// Released under an MIT-style license. See LICENSE.

package cell

import (
	"fmt"
	"math/big"
	"strconv"
)

type Atom interface {
	Cell

	Float() float64
	Int() int64
	Rat() *big.Rat
	Status() int64
}

type Cell interface {
	Bool() bool
	Equal(c Cell) bool
	String() string
}

type Number interface {
	Atom

	Greater(c Cell) bool
	Less(c Cell) bool

	Add(c Cell) Number
	Divide(c Cell) Number
	Modulo(c Cell) Number
	Multiply(c Cell) Number
	Subtract(c Cell) Number
}

type Reference interface {
	Cell

	Copy() Reference
	Get() Cell
	Set(c Cell)
}

var (
	Null  Cell
	False *Boolean
	True  *Boolean

	max *big.Rat
	min *big.Rat
	num [512]*Integer
	one *big.Rat
	rat [512]Rational
	res [256]*Status
	sym map[string]*Symbol
	zip *big.Rat
)

func init() {
	max = big.NewRat(255, 1)
	min = big.NewRat(-256, 1)

	one = big.NewRat(1, 1)
	zip = big.NewRat(0, 1)
	rat[257] = Rational{one}
	rat[256] = Rational{zip}

	pair := new(Pair)
	pair.car = pair
	pair.cdr = pair

	Null = Cell(pair)

	F := Boolean(false)
	False = &F

	T := Boolean(true)
	True = &T

	sym = make(map[string]*Symbol)

	for _, v := range [...]string{
		"$redirect",
		"append-stderr",
		"append-stdout",
		"channel-stderr",
		"channel-stdout",
		"eval-list",
		"is-boolean",
		"is-builtin",
		"is-channel",
		"is-integer",
		"is-method",
		"is-number",
		"is-object",
		"is-status",
		"is-string",
		"is-symbol",
		"is-syntax",
		"list-to-string",
		"list-to-symbol",
		"pipe-stderr",
		"pipe-stdout",
		"reader-close",
		"redirect-stderr",
		"redirect-stdin",
		"redirect-stdout",
		"substitute-stdin",
		"substitute-stdout",
		"substring",
		"writer-close",
		"is-control",
		"is-graphic",
		"is-letter",
	} {
		sym[v] = NewSymbol(v)
	}
}

func ratmod(x, y *big.Rat) *big.Rat {
	if x.IsInt() && y.IsInt() {
		return new(big.Rat).SetInt(new(big.Int).Mod(x.Num(), y.Num()))
	}

	panic("operation not permitted")
}

func AppendTo(list Cell, elements ...Cell) Cell {
	var pair, prev, start Cell

	index := 0

	start = Null

	if list == nil {
		panic("cannot append to non-existent list")
	}

	if list != Null {
		start = list

		for prev = list; Cdr(prev) != Null; prev = Cdr(prev) {
		}

	} else if len(elements) > 0 {
		start = Cons(elements[index], Null)
		prev = start
		index++
	}

	for ; index < len(elements); index++ {
		pair = Cons(elements[index], Null)
		SetCdr(prev, pair)
		prev = pair
	}

	return start
}

func Car(c Cell) Cell {
	return c.(*Pair).car
}

func Cdr(c Cell) Cell {
	return c.(*Pair).cdr
}

func Caar(c Cell) Cell {
	return c.(*Pair).car.(*Pair).car
}

func Cadr(c Cell) Cell {
	return c.(*Pair).cdr.(*Pair).car
}

func Cdar(c Cell) Cell {
	return c.(*Pair).car.(*Pair).cdr
}

func Cddr(c Cell) Cell {
	return c.(*Pair).cdr.(*Pair).cdr
}

func Caaar(c Cell) Cell {
	return c.(*Pair).car.(*Pair).car.(*Pair).car
}

func Caadr(c Cell) Cell {
	return c.(*Pair).cdr.(*Pair).car.(*Pair).car
}

func Cadar(c Cell) Cell {
	return c.(*Pair).car.(*Pair).cdr.(*Pair).car
}

func Caddr(c Cell) Cell {
	return c.(*Pair).cdr.(*Pair).cdr.(*Pair).car
}

func Cdaar(c Cell) Cell {
	return c.(*Pair).car.(*Pair).car.(*Pair).cdr
}

func Cdadr(c Cell) Cell {
	return c.(*Pair).cdr.(*Pair).car.(*Pair).cdr
}

func Cddar(c Cell) Cell {
	return c.(*Pair).car.(*Pair).cdr.(*Pair).cdr
}

func Cdddr(c Cell) Cell {
	return c.(*Pair).cdr.(*Pair).cdr.(*Pair).cdr
}

func IsAtom(c Cell) bool {
	switch c.(type) {
	case Atom:
		return true
	}
	return false
}

func IsNull(c Cell) bool {
	return c == Null
}

func IsNumber(c Cell) bool {
	switch t := c.(type) {
	case Number:
		return true
	case *Symbol:
		return t.isNumeric()
	}
	return false
}

func Join(list Cell, elements ...Cell) Cell {
	var pair, prev, start Cell

	if list != nil && list != Null {
		start = Cons(Car(list), Null)

		for list = Cdr(list); list != Null; list = Cdr(list) {
			pair = Cons(Car(list), Null)
			SetCdr(prev, pair)
			prev = pair
		}
	} else if len(elements) > 0 {
		return Join(elements[0], elements[1:]...)
	} else {
		return Null
	}

	for index := 0; index < len(elements); index++ {
		for list = elements[index]; list != Null; list = Cdr(list) {
			pair = Cons(Car(list), Null)
			SetCdr(prev, pair)
			prev = pair
		}
	}

	return start
}

func JoinTo(list Cell, elements ...Cell) Cell {
	var pair, prev, start Cell

	start = list

	if list == nil {
		panic("cannot append to non-existent list")
	} else if list == Null {
		panic("cannot destructively modify nil value")
	}

	for ; list != Null; list = Cdr(list) {
		prev = list
	}

	for index := 0; index < len(elements); index++ {
		for list = elements[index]; list != Null; list = Cdr(list) {
			pair = Cons(Car(list), Null)
			SetCdr(prev, pair)
			prev = pair
		}
	}

	return start
}

func Length(list Cell) int64 {
	var length int64

	for ; list != nil && list != Null; list = Cdr(list) {
		length++
	}

	return length
}

func List(elements ...Cell) Cell {
	if len(elements) <= 0 {
		return Null
	}

	start := Cons(elements[0], Null)
	prev := start

	for index := 1; index < len(elements); index++ {
		pair := Cons(elements[index], Null)
		SetCdr(prev, pair)
		prev = pair
	}

	return start
}

func Reverse(list Cell) Cell {
	reversed := Null

	for ; list != nil && list != Null; list = Cdr(list) {
		reversed = Cons(Car(list), reversed)
	}

	return reversed
}

func SetCar(c, value Cell) {
	c.(*Pair).car = value
}

func SetCdr(c, value Cell) {
	c.(*Pair).cdr = value
}

/* Boolean cell definition. */

type Boolean bool

func IsBoolean(c Cell) bool {
	switch c.(type) {
	case *Boolean:
		return true
	}
	return false
}

func NewBoolean(v bool) *Boolean {
	if v {
		return True
	}
	return False
}

func (b *Boolean) Bool() bool {
	return b == True
}

func (b *Boolean) Equal(c Cell) bool {
	if a, ok := c.(Atom); ok {
		return bool(*b) == a.Bool()
	}
	return false
}

func (b *Boolean) String() string {
	if b == True {
		return "true"
	}
	return "false"
}

func (b *Boolean) Float() float64 {
	if b == True {
		return 1.0
	}
	return 0.0
}

func (b *Boolean) Int() int64 {
	if b == True {
		return 1
	}
	return 0
}

func (b *Boolean) Rat() *big.Rat {
	if b == True {
		return one
	}
	return zip
}

func (b *Boolean) Status() int64 {
	if b == True {
		return 0
	}
	return 1
}

/* Constant cell definition. */

type Constant struct {
	Variable
}

func NewConstant(v Cell) *Constant {
	return &Constant{Variable{v}}
}

func (ct *Constant) String() string {
	return fmt.Sprintf("%%ct %p%%", ct)
}

func (ct *Constant) Set(c Cell) {
	panic("constant cannot be set")
}

/* Float cell definition. */

type Float float64

func IsFloat(c Cell) bool {
	switch c.(type) {
	case *Float:
		return true
	}
	return false
}

func NewFloat(v float64) *Float {
	f := Float(v)
	return &f
}

func (f *Float) Bool() bool {
	return *f != 0
}

func (f *Float) Equal(c Cell) bool {
	if a, ok := c.(Atom); ok {
		return f.Rat().Cmp(a.Rat()) == 0
	}
	return false
}

func (f *Float) String() string {
	return strconv.FormatFloat(float64(*f), 'g', -1, 64)
}

func (f *Float) Float() float64 {
	return float64(*f)
}

func (f *Float) Int() int64 {
	return int64(*f)
}

func (f *Float) Rat() *big.Rat {
	return new(big.Rat).SetFloat64(float64(*f))
}

func (f *Float) Status() int64 {
	return int64(*f)
}

func (f *Float) Greater(c Cell) bool {
	return f.Rat().Cmp(c.(Atom).Rat()) > 0
}

func (f *Float) Less(c Cell) bool {
	return f.Rat().Cmp(c.(Atom).Rat()) < 0
}

func (f *Float) Add(c Cell) Number {
	return NewRational(new(big.Rat).Add(f.Rat(), c.(Atom).Rat()))
}

func (f *Float) Divide(c Cell) Number {
	return NewRational(new(big.Rat).Quo(f.Rat(), c.(Atom).Rat()))
}

func (f *Float) Modulo(c Cell) Number {
	return NewRational(ratmod(f.Rat(), c.(Atom).Rat()))
}

func (f *Float) Multiply(c Cell) Number {
	return NewRational(new(big.Rat).Mul(f.Rat(), c.(Atom).Rat()))
}

func (f *Float) Subtract(c Cell) Number {
	return NewRational(new(big.Rat).Sub(f.Rat(), c.(Atom).Rat()))
}

/* Integer cell definition. */

type Integer int64

func IsInteger(c Cell) bool {
	switch c.(type) {
	case *Integer:
		return true
	}
	return false
}

func NewInteger(v int64) *Integer {
	if -256 <= v && v <= 255 {
		n := v + 256
		p := num[n]

		if p == nil {
			i := Integer(v)
			p = &i

			num[n] = p
		}

		return p
	}

	i := Integer(v)
	return &i
}

func (i *Integer) Bool() bool {
	return *i != 0
}

func (i *Integer) Equal(c Cell) bool {
	if a, ok := c.(Atom); ok {
		return i.Rat().Cmp(a.Rat()) == 0
	}
	return false
}

func (i *Integer) String() string {
	return strconv.FormatInt(int64(*i), 10)
}

func (i *Integer) Float() float64 {
	return float64(*i)
}

func (i *Integer) Int() int64 {
	return int64(*i)
}

func (i *Integer) Rat() *big.Rat {
	return big.NewRat(int64(*i), 1)
}

func (i *Integer) Status() int64 {
	return int64(*i)
}

func (i *Integer) Greater(c Cell) bool {
	return i.Rat().Cmp(c.(Atom).Rat()) > 0
}

func (i *Integer) Less(c Cell) bool {
	return i.Rat().Cmp(c.(Atom).Rat()) < 0
}

func (i *Integer) Add(c Cell) Number {
	return NewRational(new(big.Rat).Add(i.Rat(), c.(Atom).Rat()))
}

func (i *Integer) Divide(c Cell) Number {
	return NewRational(new(big.Rat).Quo(i.Rat(), c.(Atom).Rat()))
}

func (i *Integer) Modulo(c Cell) Number {
	return NewRational(ratmod(i.Rat(), c.(Atom).Rat()))
}

func (i *Integer) Multiply(c Cell) Number {
	return NewRational(new(big.Rat).Mul(i.Rat(), c.(Atom).Rat()))
}

func (i *Integer) Subtract(c Cell) Number {
	return NewRational(new(big.Rat).Sub(i.Rat(), c.(Atom).Rat()))
}

/* Pair cell definition. */

type Pair struct {
	car Cell
	cdr Cell
}

func IsCons(c Cell) bool {
	switch c.(type) {
	case *Pair:
		return true
	}
	return false
}

func Cons(h, t Cell) Cell {
	return &Pair{car: h, cdr: t}
}

func (p *Pair) Bool() bool {
	return p != Null
}

func (p *Pair) Equal(c Cell) bool {
	if p == Null && c == Null {
		return true
	}
	return p.car.Equal(Car(c)) && p.cdr.Equal(Cdr(c))
}

func (p *Pair) String() (s string) {
	s = ""

	if IsCons(p.car) && IsCons(Cdr(p.car)) {
		s += "("
	}

	if p.car != Null {
		s += p.car.String()
	}

	if IsCons(p.car) && IsCons(Cdr(p.car)) {
		s += ")"
	}

	if IsCons(p.cdr) {
		if p.cdr == Null {
			return s
		}

		s += " "
	} else {
		s += "::"
	}

	s += p.cdr.String()

	return s
}

/* Rational cell definition. */

type Rational struct {
	v *big.Rat
}

func IsRational(c Cell) bool {
	switch c.(type) {
	case Rational:
		return true
	}
	return false
}

func NewRational(r *big.Rat) Rational {
	if !r.IsInt() || r.Cmp(min) < 0 || r.Cmp(max) > 0 {
		return Rational{r}
	}

	n := r.Num().Int64()
	i := n + 256
	p := rat[i]

	if p.v == nil {
		p = Rational{r}
		rat[i] = p
	}

	return p
}

func (r Rational) Bool() bool {
	return r.v.Cmp(zip) != 0
}

func (r Rational) Equal(c Cell) bool {
	if a, ok := c.(Atom); ok {
		return r.v.Cmp(a.Rat()) == 0
	}
	return false
}

func (r Rational) String() string {
	return r.v.RatString()
}

func (r Rational) Float() float64 {
	f, _ := r.v.Float64()
	return f
}

func (r Rational) Int() int64 {
	n := r.v.Num()
	d := r.v.Denom()
	return new(big.Int).Div(n, d).Int64()
}

func (r Rational) Rat() *big.Rat {
	return r.v
}

func (r Rational) Status() int64 {
	return r.Int()
}

func (r Rational) Greater(c Cell) bool {
	return r.v.Cmp(c.(Atom).Rat()) > 0
}

func (r Rational) Less(c Cell) bool {
	return r.v.Cmp(c.(Atom).Rat()) < 0
}

func (r Rational) Add(c Cell) Number {
	return NewRational(new(big.Rat).Add(r.v, c.(Atom).Rat()))
}

func (r Rational) Divide(c Cell) Number {
	return NewRational(new(big.Rat).Quo(r.v, c.(Atom).Rat()))
}

func (r Rational) Modulo(c Cell) Number {
	return NewRational(ratmod(r.v, c.(Atom).Rat()))
}

func (r Rational) Multiply(c Cell) Number {
	return NewRational(new(big.Rat).Mul(r.v, c.(Atom).Rat()))
}

func (r Rational) Subtract(c Cell) Number {
	return NewRational(new(big.Rat).Sub(r.v, c.(Atom).Rat()))
}

/* Status cell definition. */

type Status int64

func IsStatus(c Cell) bool {
	switch c.(type) {
	case *Status:
		return true
	}
	return false
}

func NewStatus(v int64) *Status {
	if 0 <= v && v <= 255 {
		p := res[v]

		if p == nil {
			s := Status(v)
			p = &s

			res[v] = p
		}

		return p
	}

	s := Status(v)
	return &s
}

func (s *Status) Bool() bool {
	return int64(*s) == 0
}

func (s *Status) Equal(c Cell) bool {
	if a, ok := c.(Atom); ok {
		return int64(*s) == a.Status()
	}
	return false
}

func (s *Status) String() string {
	return strconv.FormatInt(int64(*s), 10)
}

func (s *Status) Float() float64 {
	return float64(*s)
}

func (s *Status) Int() int64 {
	return int64(*s)
}

func (s *Status) Rat() *big.Rat {
	return big.NewRat(int64(*s), 1)
}

func (s *Status) Status() int64 {
	return int64(*s)
}

func (s *Status) Greater(c Cell) bool {
	return s.Rat().Cmp(c.(Atom).Rat()) > 0
}

func (s *Status) Less(c Cell) bool {
	return s.Rat().Cmp(c.(Atom).Rat()) < 0
}

func (s *Status) Add(c Cell) Number {
	return NewRational(new(big.Rat).Add(s.Rat(), c.(Atom).Rat()))
}

func (s *Status) Divide(c Cell) Number {
	return NewRational(new(big.Rat).Quo(s.Rat(), c.(Atom).Rat()))
}

func (s *Status) Modulo(c Cell) Number {
	return NewRational(ratmod(s.Rat(), c.(Atom).Rat()))
}

func (s *Status) Multiply(c Cell) Number {
	return NewRational(new(big.Rat).Mul(s.Rat(), c.(Atom).Rat()))
}

func (s *Status) Subtract(c Cell) Number {
	return NewRational(new(big.Rat).Sub(s.Rat(), c.(Atom).Rat()))
}

/* Symbol cell definition. */

type Symbol string

func IsSymbol(c Cell) bool {
	switch c.(type) {
	case *Symbol:
		return true
	}
	return false
}

func NewSymbol(v string) *Symbol {
	p, ok := sym[v]

	if ok {
		return p
	}

	s := Symbol(v)
	p = &s

	if len(v) <= 8 {
		sym[v] = p
	}

	return p
}

func (s *Symbol) Bool() bool {
	if string(*s) == "false" {
		return false
	}

	return true
}

func (s *Symbol) Equal(c Cell) bool {
	if a, ok := c.(Atom); ok {
		return string(*s) == a.String()
	}
	return false
}

func (s *Symbol) String() string {
	return string(*s)
}

func (s *Symbol) Float() (f float64) {
	var err error
	if f, err = strconv.ParseFloat(string(*s), 64); err != nil {
		panic(err)
	}
	return f
}

func (s *Symbol) Int() (i int64) {
	var err error
	if i, err = strconv.ParseInt(string(*s), 0, 64); err != nil {
		panic(err)
	}
	return i
}

func (s *Symbol) Rat() *big.Rat {
	r := new(big.Rat)
	if _, err := fmt.Sscan(string(*s), r); err != nil {
		panic(err)
	}
	return r
}

func (s *Symbol) Status() (i int64) {
	var err error
	if i, err = strconv.ParseInt(string(*s), 0, 64); err != nil {
		panic(err)
	}
	return i
}

func (s *Symbol) Greater(c Cell) bool {
	return string(*s) > c.(Atom).String()
}

func (s *Symbol) Less(c Cell) bool {
	return string(*s) < c.(Atom).String()
}

func (s *Symbol) Add(c Cell) Number {
	return NewRational(new(big.Rat).Add(s.Rat(), c.(Atom).Rat()))
}

func (s *Symbol) Divide(c Cell) Number {
	return NewRational(new(big.Rat).Quo(s.Rat(), c.(Atom).Rat()))
}

func (s *Symbol) Modulo(c Cell) Number {
	return NewRational(ratmod(s.Rat(), c.(Atom).Rat()))
}

func (s *Symbol) Multiply(c Cell) Number {
	return NewRational(new(big.Rat).Mul(s.Rat(), c.(Atom).Rat()))
}

func (s *Symbol) Subtract(c Cell) Number {
	return NewRational(new(big.Rat).Sub(s.Rat(), c.(Atom).Rat()))
}

/* Symbol-specific functions. */

func (s *Symbol) isNumeric() bool {
	r := new(big.Rat)
	_, err := fmt.Sscan(string(*s), r)
	return err == nil
}

/* Variable cell definition. */

type Variable struct {
	v Cell
}

func NewVariable(v Cell) Reference {
	return &Variable{v}
}

func (vr *Variable) Bool() bool {
	return true
}

func (vr *Variable) Equal(c Cell) bool {
	return vr.v.Equal(c)
}

func (vr *Variable) String() string {
	return fmt.Sprintf("%%variable %p%%", vr)
}

/* Variable-specific functions */

func (vr *Variable) Copy() Reference {
	return NewVariable(vr.v)
}

func (vr *Variable) Get() Cell {
	return vr.v
}

func (vr *Variable) Set(c Cell) {
	vr.v = c
}
