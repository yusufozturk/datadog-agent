// Code generated - DO NOT EDIT.

package eval

type StringOpOverloadBase interface {
	StringEquals(ctx *Context, value string) bool

	StringNotEquals(ctx *Context, value string) bool
}

type IntOpOverloadBase interface {
	IntEquals(ctx *Context, value int) bool

	IntNotEquals(ctx *Context, value int) bool

	IntAnd(ctx *Context, value int) int

	IntOr(ctx *Context, value int) int

	IntXor(ctx *Context, value int) int

	GreaterThan(ctx *Context, value int) bool

	GreaterOrEqualThan(ctx *Context, value int) bool

	LesserThan(ctx *Context, value int) bool

	LesserOrEqualThan(ctx *Context, value int) bool
}

type BoolOpOverloadBase interface {
	BoolEquals(ctx *Context, value bool) bool

	BoolNotEquals(ctx *Context, value bool) bool
}

func Or(a *BoolEvaluator, b *BoolEvaluator, opts *Opts, state *state) (*BoolEvaluator, error) {
	partialA, partialB := a.isPartial, b.isPartial

	if a.EvalFnc == nil || (a.Field != "" && a.Field != state.field) {
		partialA = true
	}
	if b.EvalFnc == nil || (b.Field != "" && b.Field != state.field) {
		partialB = true
	}
	isPartialLeaf := partialA && partialB

	if a.Field != "" && b.Field != "" {
		isPartialLeaf = true
	}

	if (a.EvalFnc != nil || a.OpOverload != nil) && (b.EvalFnc != nil || b.OpOverload != nil) {
		ea, eb := a.EvalFnc, b.EvalFnc

		if state.field != "" {
			if a.isPartial {
				ea = func(ctx *Context) bool {
					return true
				}
			}
			if b.isPartial {
				eb = func(ctx *Context) bool {
					return true
				}
			}
		}

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb(ctx)
				result := op1 || op2
				ctx.Logf("Evaluating %v || %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			evalFnc = func(ctx *Context) bool {
				return ea(ctx) || eb(ctx)
			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc == nil && a.OpOverload == nil && b.EvalFnc == nil && b.OpOverload == nil {
		ea, eb := a.Value, b.Value

		if state.field != "" {
			if a.isPartial {
				ea = true
			}
			if b.isPartial {
				eb = true
			}
		}

		return &BoolEvaluator{
			Value:     ea || eb,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc != nil || a.OpOverload != nil {
		ea, eb := a.EvalFnc, b.Value

		if a.Field != "" {
			if err := state.UpdateFieldValues(a.Field, FieldValue{Value: eb, Type: ScalarValueType}); err != nil {
				return nil, err
			}
		}

		if state.field != "" {
			if a.isPartial {
				ea = func(ctx *Context) bool {
					return true
				}
			}
			if b.isPartial {
				eb = true
			}
		}

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb
				result := op1 || op2
				ctx.Logf("Evaluating %v || %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			evalFnc = func(ctx *Context) bool {
				return ea(ctx) || eb
			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	ea, eb := a.Value, b.EvalFnc

	if b.Field != "" {
		if err := state.UpdateFieldValues(b.Field, FieldValue{Value: ea, Type: ScalarValueType}); err != nil {
			return nil, err
		}
	}

	if state.field != "" {
		if a.isPartial {
			ea = true
		}
		if b.isPartial {
			eb = func(ctx *Context) bool {
				return true
			}
		}
	}

	var evalFnc func(ctx *Context) bool
	if opts.Debug {
		evalFnc = func(ctx *Context) bool {
			ctx.evalDepth++
			op1, op2 := ea, eb(ctx)
			result := op1 || op2
			ctx.Logf("Evaluating %v || %v => %v", op1, op2, result)
			ctx.evalDepth--
			return result
		}
	} else {

		evalFnc = func(ctx *Context) bool {
			return ea || eb(ctx)
		}

	}

	return &BoolEvaluator{
		EvalFnc:   evalFnc,
		isPartial: isPartialLeaf,
	}, nil
}

func And(a *BoolEvaluator, b *BoolEvaluator, opts *Opts, state *state) (*BoolEvaluator, error) {
	partialA, partialB := a.isPartial, b.isPartial

	if a.EvalFnc == nil || (a.Field != "" && a.Field != state.field) {
		partialA = true
	}
	if b.EvalFnc == nil || (b.Field != "" && b.Field != state.field) {
		partialB = true
	}
	isPartialLeaf := partialA && partialB

	if a.Field != "" && b.Field != "" {
		isPartialLeaf = true
	}

	if (a.EvalFnc != nil || a.OpOverload != nil) && (b.EvalFnc != nil || b.OpOverload != nil) {
		ea, eb := a.EvalFnc, b.EvalFnc

		if state.field != "" {
			if a.isPartial {
				ea = func(ctx *Context) bool {
					return true
				}
			}
			if b.isPartial {
				eb = func(ctx *Context) bool {
					return true
				}
			}
		}

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb(ctx)
				result := op1 && op2
				ctx.Logf("Evaluating %v && %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			evalFnc = func(ctx *Context) bool {
				return ea(ctx) && eb(ctx)
			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc == nil && a.OpOverload == nil && b.EvalFnc == nil && b.OpOverload == nil {
		ea, eb := a.Value, b.Value

		if state.field != "" {
			if a.isPartial {
				ea = true
			}
			if b.isPartial {
				eb = true
			}
		}

		return &BoolEvaluator{
			Value:     ea && eb,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc != nil || a.OpOverload != nil {
		ea, eb := a.EvalFnc, b.Value

		if a.Field != "" {
			if err := state.UpdateFieldValues(a.Field, FieldValue{Value: eb, Type: ScalarValueType}); err != nil {
				return nil, err
			}
		}

		if state.field != "" {
			if a.isPartial {
				ea = func(ctx *Context) bool {
					return true
				}
			}
			if b.isPartial {
				eb = true
			}
		}

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb
				result := op1 && op2
				ctx.Logf("Evaluating %v && %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			evalFnc = func(ctx *Context) bool {
				return ea(ctx) && eb
			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	ea, eb := a.Value, b.EvalFnc

	if b.Field != "" {
		if err := state.UpdateFieldValues(b.Field, FieldValue{Value: ea, Type: ScalarValueType}); err != nil {
			return nil, err
		}
	}

	if state.field != "" {
		if a.isPartial {
			ea = true
		}
		if b.isPartial {
			eb = func(ctx *Context) bool {
				return true
			}
		}
	}

	var evalFnc func(ctx *Context) bool
	if opts.Debug {
		evalFnc = func(ctx *Context) bool {
			ctx.evalDepth++
			op1, op2 := ea, eb(ctx)
			result := op1 && op2
			ctx.Logf("Evaluating %v && %v => %v", op1, op2, result)
			ctx.evalDepth--
			return result
		}
	} else {

		evalFnc = func(ctx *Context) bool {
			return ea && eb(ctx)
		}

	}

	return &BoolEvaluator{
		EvalFnc:   evalFnc,
		isPartial: isPartialLeaf,
	}, nil
}

func IntEquals(a *IntEvaluator, b *IntEvaluator, opts *Opts, state *state) (*BoolEvaluator, error) {
	partialA, partialB := a.isPartial, b.isPartial

	if a.EvalFnc == nil || (a.Field != "" && a.Field != state.field) {
		partialA = true
	}
	if b.EvalFnc == nil || (b.Field != "" && b.Field != state.field) {
		partialB = true
	}
	isPartialLeaf := partialA && partialB

	if a.Field != "" && b.Field != "" {
		isPartialLeaf = true
	}

	if (a.EvalFnc != nil || a.OpOverload != nil) && (b.EvalFnc != nil || b.OpOverload != nil) {
		ea, eb := a.EvalFnc, b.EvalFnc

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb(ctx)
				result := op1 == op2
				ctx.Logf("Evaluating %v == %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return a.OpOverload.IntEquals(ctx, eb(ctx))
				}
			} else if b.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return b.OpOverload.IntEquals(ctx, ea(ctx))
				}
			} else {

				evalFnc = func(ctx *Context) bool {
					return ea(ctx) == eb(ctx)
				}

			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc == nil && a.OpOverload == nil && b.EvalFnc == nil && b.OpOverload == nil {
		ea, eb := a.Value, b.Value

		return &BoolEvaluator{
			Value:     ea == eb,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc != nil || a.OpOverload != nil {
		ea, eb := a.EvalFnc, b.Value

		if a.Field != "" {
			if err := state.UpdateFieldValues(a.Field, FieldValue{Value: eb, Type: ScalarValueType}); err != nil {
				return nil, err
			}
		}

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb
				result := op1 == op2
				ctx.Logf("Evaluating %v == %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return a.OpOverload.IntEquals(ctx, eb)
				}
			} else {

				evalFnc = func(ctx *Context) bool {
					return ea(ctx) == eb
				}

			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	ea, eb := a.Value, b.EvalFnc

	if b.Field != "" {
		if err := state.UpdateFieldValues(b.Field, FieldValue{Value: ea, Type: ScalarValueType}); err != nil {
			return nil, err
		}
	}

	var evalFnc func(ctx *Context) bool
	if opts.Debug {
		evalFnc = func(ctx *Context) bool {
			ctx.evalDepth++
			op1, op2 := ea, eb(ctx)
			result := op1 == op2
			ctx.Logf("Evaluating %v == %v => %v", op1, op2, result)
			ctx.evalDepth--
			return result
		}
	} else {

		if a.OpOverload != nil {
			evalFnc = func(ctx *Context) bool {
				return b.OpOverload.IntEquals(ctx, ea)
			}
		} else {

			evalFnc = func(ctx *Context) bool {
				return ea == eb(ctx)
			}

		}

	}

	return &BoolEvaluator{
		EvalFnc:   evalFnc,
		isPartial: isPartialLeaf,
	}, nil
}

func IntNotEquals(a *IntEvaluator, b *IntEvaluator, opts *Opts, state *state) (*BoolEvaluator, error) {
	partialA, partialB := a.isPartial, b.isPartial

	if a.EvalFnc == nil || (a.Field != "" && a.Field != state.field) {
		partialA = true
	}
	if b.EvalFnc == nil || (b.Field != "" && b.Field != state.field) {
		partialB = true
	}
	isPartialLeaf := partialA && partialB

	if a.Field != "" && b.Field != "" {
		isPartialLeaf = true
	}

	if (a.EvalFnc != nil || a.OpOverload != nil) && (b.EvalFnc != nil || b.OpOverload != nil) {
		ea, eb := a.EvalFnc, b.EvalFnc

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb(ctx)
				result := op1 != op2
				ctx.Logf("Evaluating %v != %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return a.OpOverload.IntNotEquals(ctx, eb(ctx))
				}
			} else if b.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return b.OpOverload.IntNotEquals(ctx, ea(ctx))
				}
			} else {

				evalFnc = func(ctx *Context) bool {
					return ea(ctx) != eb(ctx)
				}

			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc == nil && a.OpOverload == nil && b.EvalFnc == nil && b.OpOverload == nil {
		ea, eb := a.Value, b.Value

		return &BoolEvaluator{
			Value:     ea != eb,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc != nil || a.OpOverload != nil {
		ea, eb := a.EvalFnc, b.Value

		if a.Field != "" {
			if err := state.UpdateFieldValues(a.Field, FieldValue{Value: eb, Type: ScalarValueType}); err != nil {
				return nil, err
			}
		}

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb
				result := op1 != op2
				ctx.Logf("Evaluating %v != %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return a.OpOverload.IntNotEquals(ctx, eb)
				}
			} else {

				evalFnc = func(ctx *Context) bool {
					return ea(ctx) != eb
				}

			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	ea, eb := a.Value, b.EvalFnc

	if b.Field != "" {
		if err := state.UpdateFieldValues(b.Field, FieldValue{Value: ea, Type: ScalarValueType}); err != nil {
			return nil, err
		}
	}

	var evalFnc func(ctx *Context) bool
	if opts.Debug {
		evalFnc = func(ctx *Context) bool {
			ctx.evalDepth++
			op1, op2 := ea, eb(ctx)
			result := op1 != op2
			ctx.Logf("Evaluating %v != %v => %v", op1, op2, result)
			ctx.evalDepth--
			return result
		}
	} else {

		if a.OpOverload != nil {
			evalFnc = func(ctx *Context) bool {
				return b.OpOverload.IntNotEquals(ctx, ea)
			}
		} else {

			evalFnc = func(ctx *Context) bool {
				return ea != eb(ctx)
			}

		}

	}

	return &BoolEvaluator{
		EvalFnc:   evalFnc,
		isPartial: isPartialLeaf,
	}, nil
}

func IntAnd(a *IntEvaluator, b *IntEvaluator, opts *Opts, state *state) (*IntEvaluator, error) {
	partialA, partialB := a.isPartial, b.isPartial

	if a.EvalFnc == nil || (a.Field != "" && a.Field != state.field) {
		partialA = true
	}
	if b.EvalFnc == nil || (b.Field != "" && b.Field != state.field) {
		partialB = true
	}
	isPartialLeaf := partialA && partialB

	if a.Field != "" && b.Field != "" {
		isPartialLeaf = true
	}

	if (a.EvalFnc != nil || a.OpOverload != nil) && (b.EvalFnc != nil || b.OpOverload != nil) {
		ea, eb := a.EvalFnc, b.EvalFnc

		var evalFnc func(ctx *Context) int
		if opts.Debug {
			evalFnc = func(ctx *Context) int {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb(ctx)
				result := op1 & op2
				ctx.Logf("Evaluating %v & %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) int {
					return a.OpOverload.IntAnd(ctx, eb(ctx))
				}
			} else if b.OpOverload != nil {
				evalFnc = func(ctx *Context) int {
					return b.OpOverload.IntAnd(ctx, ea(ctx))
				}
			} else {

				evalFnc = func(ctx *Context) int {
					return ea(ctx) & eb(ctx)
				}

			}

		}

		return &IntEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc == nil && a.OpOverload == nil && b.EvalFnc == nil && b.OpOverload == nil {
		ea, eb := a.Value, b.Value

		return &IntEvaluator{
			Value:     ea & eb,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc != nil || a.OpOverload != nil {
		ea, eb := a.EvalFnc, b.Value

		if a.Field != "" {
			if err := state.UpdateFieldValues(a.Field, FieldValue{Value: eb, Type: BitmaskValueType}); err != nil {
				return nil, err
			}
		}

		var evalFnc func(ctx *Context) int
		if opts.Debug {
			evalFnc = func(ctx *Context) int {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb
				result := op1 & op2
				ctx.Logf("Evaluating %v & %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) int {
					return a.OpOverload.IntAnd(ctx, eb)
				}
			} else {

				evalFnc = func(ctx *Context) int {
					return ea(ctx) & eb
				}

			}

		}

		return &IntEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	ea, eb := a.Value, b.EvalFnc

	if b.Field != "" {
		if err := state.UpdateFieldValues(b.Field, FieldValue{Value: ea, Type: BitmaskValueType}); err != nil {
			return nil, err
		}
	}

	var evalFnc func(ctx *Context) int
	if opts.Debug {
		evalFnc = func(ctx *Context) int {
			ctx.evalDepth++
			op1, op2 := ea, eb(ctx)
			result := op1 & op2
			ctx.Logf("Evaluating %v & %v => %v", op1, op2, result)
			ctx.evalDepth--
			return result
		}
	} else {

		if a.OpOverload != nil {
			evalFnc = func(ctx *Context) int {
				return b.OpOverload.IntAnd(ctx, ea)
			}
		} else {

			evalFnc = func(ctx *Context) int {
				return ea & eb(ctx)
			}

		}

	}

	return &IntEvaluator{
		EvalFnc:   evalFnc,
		isPartial: isPartialLeaf,
	}, nil
}

func IntOr(a *IntEvaluator, b *IntEvaluator, opts *Opts, state *state) (*IntEvaluator, error) {
	partialA, partialB := a.isPartial, b.isPartial

	if a.EvalFnc == nil || (a.Field != "" && a.Field != state.field) {
		partialA = true
	}
	if b.EvalFnc == nil || (b.Field != "" && b.Field != state.field) {
		partialB = true
	}
	isPartialLeaf := partialA && partialB

	if a.Field != "" && b.Field != "" {
		isPartialLeaf = true
	}

	if (a.EvalFnc != nil || a.OpOverload != nil) && (b.EvalFnc != nil || b.OpOverload != nil) {
		ea, eb := a.EvalFnc, b.EvalFnc

		var evalFnc func(ctx *Context) int
		if opts.Debug {
			evalFnc = func(ctx *Context) int {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb(ctx)
				result := op1 | op2
				ctx.Logf("Evaluating %v | %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) int {
					return a.OpOverload.IntOr(ctx, eb(ctx))
				}
			} else if b.OpOverload != nil {
				evalFnc = func(ctx *Context) int {
					return b.OpOverload.IntOr(ctx, ea(ctx))
				}
			} else {

				evalFnc = func(ctx *Context) int {
					return ea(ctx) | eb(ctx)
				}

			}

		}

		return &IntEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc == nil && a.OpOverload == nil && b.EvalFnc == nil && b.OpOverload == nil {
		ea, eb := a.Value, b.Value

		return &IntEvaluator{
			Value:     ea | eb,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc != nil || a.OpOverload != nil {
		ea, eb := a.EvalFnc, b.Value

		if a.Field != "" {
			if err := state.UpdateFieldValues(a.Field, FieldValue{Value: eb, Type: BitmaskValueType}); err != nil {
				return nil, err
			}
		}

		var evalFnc func(ctx *Context) int
		if opts.Debug {
			evalFnc = func(ctx *Context) int {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb
				result := op1 | op2
				ctx.Logf("Evaluating %v | %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) int {
					return a.OpOverload.IntOr(ctx, eb)
				}
			} else {

				evalFnc = func(ctx *Context) int {
					return ea(ctx) | eb
				}

			}

		}

		return &IntEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	ea, eb := a.Value, b.EvalFnc

	if b.Field != "" {
		if err := state.UpdateFieldValues(b.Field, FieldValue{Value: ea, Type: BitmaskValueType}); err != nil {
			return nil, err
		}
	}

	var evalFnc func(ctx *Context) int
	if opts.Debug {
		evalFnc = func(ctx *Context) int {
			ctx.evalDepth++
			op1, op2 := ea, eb(ctx)
			result := op1 | op2
			ctx.Logf("Evaluating %v | %v => %v", op1, op2, result)
			ctx.evalDepth--
			return result
		}
	} else {

		if a.OpOverload != nil {
			evalFnc = func(ctx *Context) int {
				return b.OpOverload.IntOr(ctx, ea)
			}
		} else {

			evalFnc = func(ctx *Context) int {
				return ea | eb(ctx)
			}

		}

	}

	return &IntEvaluator{
		EvalFnc:   evalFnc,
		isPartial: isPartialLeaf,
	}, nil
}

func IntXor(a *IntEvaluator, b *IntEvaluator, opts *Opts, state *state) (*IntEvaluator, error) {
	partialA, partialB := a.isPartial, b.isPartial

	if a.EvalFnc == nil || (a.Field != "" && a.Field != state.field) {
		partialA = true
	}
	if b.EvalFnc == nil || (b.Field != "" && b.Field != state.field) {
		partialB = true
	}
	isPartialLeaf := partialA && partialB

	if a.Field != "" && b.Field != "" {
		isPartialLeaf = true
	}

	if (a.EvalFnc != nil || a.OpOverload != nil) && (b.EvalFnc != nil || b.OpOverload != nil) {
		ea, eb := a.EvalFnc, b.EvalFnc

		var evalFnc func(ctx *Context) int
		if opts.Debug {
			evalFnc = func(ctx *Context) int {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb(ctx)
				result := op1 ^ op2
				ctx.Logf("Evaluating %v ^ %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) int {
					return a.OpOverload.IntXor(ctx, eb(ctx))
				}
			} else if b.OpOverload != nil {
				evalFnc = func(ctx *Context) int {
					return b.OpOverload.IntXor(ctx, ea(ctx))
				}
			} else {

				evalFnc = func(ctx *Context) int {
					return ea(ctx) ^ eb(ctx)
				}

			}

		}

		return &IntEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc == nil && a.OpOverload == nil && b.EvalFnc == nil && b.OpOverload == nil {
		ea, eb := a.Value, b.Value

		return &IntEvaluator{
			Value:     ea ^ eb,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc != nil || a.OpOverload != nil {
		ea, eb := a.EvalFnc, b.Value

		if a.Field != "" {
			if err := state.UpdateFieldValues(a.Field, FieldValue{Value: eb, Type: BitmaskValueType}); err != nil {
				return nil, err
			}
		}

		var evalFnc func(ctx *Context) int
		if opts.Debug {
			evalFnc = func(ctx *Context) int {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb
				result := op1 ^ op2
				ctx.Logf("Evaluating %v ^ %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) int {
					return a.OpOverload.IntXor(ctx, eb)
				}
			} else {

				evalFnc = func(ctx *Context) int {
					return ea(ctx) ^ eb
				}

			}

		}

		return &IntEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	ea, eb := a.Value, b.EvalFnc

	if b.Field != "" {
		if err := state.UpdateFieldValues(b.Field, FieldValue{Value: ea, Type: BitmaskValueType}); err != nil {
			return nil, err
		}
	}

	var evalFnc func(ctx *Context) int
	if opts.Debug {
		evalFnc = func(ctx *Context) int {
			ctx.evalDepth++
			op1, op2 := ea, eb(ctx)
			result := op1 ^ op2
			ctx.Logf("Evaluating %v ^ %v => %v", op1, op2, result)
			ctx.evalDepth--
			return result
		}
	} else {

		if a.OpOverload != nil {
			evalFnc = func(ctx *Context) int {
				return b.OpOverload.IntXor(ctx, ea)
			}
		} else {

			evalFnc = func(ctx *Context) int {
				return ea ^ eb(ctx)
			}

		}

	}

	return &IntEvaluator{
		EvalFnc:   evalFnc,
		isPartial: isPartialLeaf,
	}, nil
}

func StringEquals(a *StringEvaluator, b *StringEvaluator, opts *Opts, state *state) (*BoolEvaluator, error) {
	partialA, partialB := a.isPartial, b.isPartial

	if a.EvalFnc == nil || (a.Field != "" && a.Field != state.field) {
		partialA = true
	}
	if b.EvalFnc == nil || (b.Field != "" && b.Field != state.field) {
		partialB = true
	}
	isPartialLeaf := partialA && partialB

	if a.Field != "" && b.Field != "" {
		isPartialLeaf = true
	}

	if (a.EvalFnc != nil || a.OpOverload != nil) && (b.EvalFnc != nil || b.OpOverload != nil) {
		ea, eb := a.EvalFnc, b.EvalFnc

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb(ctx)
				result := op1 == op2
				ctx.Logf("Evaluating %v == %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return a.OpOverload.StringEquals(ctx, eb(ctx))
				}
			} else if b.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return b.OpOverload.StringEquals(ctx, ea(ctx))
				}
			} else {

				evalFnc = func(ctx *Context) bool {
					return ea(ctx) == eb(ctx)
				}

			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc == nil && a.OpOverload == nil && b.EvalFnc == nil && b.OpOverload == nil {
		ea, eb := a.Value, b.Value

		return &BoolEvaluator{
			Value:     ea == eb,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc != nil || a.OpOverload != nil {
		ea, eb := a.EvalFnc, b.Value

		if a.Field != "" {
			if err := state.UpdateFieldValues(a.Field, FieldValue{Value: eb, Type: ScalarValueType}); err != nil {
				return nil, err
			}
		}

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb
				result := op1 == op2
				ctx.Logf("Evaluating %v == %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return a.OpOverload.StringEquals(ctx, eb)
				}
			} else {

				evalFnc = func(ctx *Context) bool {
					return ea(ctx) == eb
				}

			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	ea, eb := a.Value, b.EvalFnc

	if b.Field != "" {
		if err := state.UpdateFieldValues(b.Field, FieldValue{Value: ea, Type: ScalarValueType}); err != nil {
			return nil, err
		}
	}

	var evalFnc func(ctx *Context) bool
	if opts.Debug {
		evalFnc = func(ctx *Context) bool {
			ctx.evalDepth++
			op1, op2 := ea, eb(ctx)
			result := op1 == op2
			ctx.Logf("Evaluating %v == %v => %v", op1, op2, result)
			ctx.evalDepth--
			return result
		}
	} else {

		if a.OpOverload != nil {
			evalFnc = func(ctx *Context) bool {
				return b.OpOverload.StringEquals(ctx, ea)
			}
		} else {

			evalFnc = func(ctx *Context) bool {
				return ea == eb(ctx)
			}

		}

	}

	return &BoolEvaluator{
		EvalFnc:   evalFnc,
		isPartial: isPartialLeaf,
	}, nil
}

func StringNotEquals(a *StringEvaluator, b *StringEvaluator, opts *Opts, state *state) (*BoolEvaluator, error) {
	partialA, partialB := a.isPartial, b.isPartial

	if a.EvalFnc == nil || (a.Field != "" && a.Field != state.field) {
		partialA = true
	}
	if b.EvalFnc == nil || (b.Field != "" && b.Field != state.field) {
		partialB = true
	}
	isPartialLeaf := partialA && partialB

	if a.Field != "" && b.Field != "" {
		isPartialLeaf = true
	}

	if (a.EvalFnc != nil || a.OpOverload != nil) && (b.EvalFnc != nil || b.OpOverload != nil) {
		ea, eb := a.EvalFnc, b.EvalFnc

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb(ctx)
				result := op1 != op2
				ctx.Logf("Evaluating %v != %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return a.OpOverload.StringNotEquals(ctx, eb(ctx))
				}
			} else if b.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return b.OpOverload.StringNotEquals(ctx, ea(ctx))
				}
			} else {

				evalFnc = func(ctx *Context) bool {
					return ea(ctx) != eb(ctx)
				}

			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc == nil && a.OpOverload == nil && b.EvalFnc == nil && b.OpOverload == nil {
		ea, eb := a.Value, b.Value

		return &BoolEvaluator{
			Value:     ea != eb,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc != nil || a.OpOverload != nil {
		ea, eb := a.EvalFnc, b.Value

		if a.Field != "" {
			if err := state.UpdateFieldValues(a.Field, FieldValue{Value: eb, Type: ScalarValueType}); err != nil {
				return nil, err
			}
		}

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb
				result := op1 != op2
				ctx.Logf("Evaluating %v != %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return a.OpOverload.StringNotEquals(ctx, eb)
				}
			} else {

				evalFnc = func(ctx *Context) bool {
					return ea(ctx) != eb
				}

			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	ea, eb := a.Value, b.EvalFnc

	if b.Field != "" {
		if err := state.UpdateFieldValues(b.Field, FieldValue{Value: ea, Type: ScalarValueType}); err != nil {
			return nil, err
		}
	}

	var evalFnc func(ctx *Context) bool
	if opts.Debug {
		evalFnc = func(ctx *Context) bool {
			ctx.evalDepth++
			op1, op2 := ea, eb(ctx)
			result := op1 != op2
			ctx.Logf("Evaluating %v != %v => %v", op1, op2, result)
			ctx.evalDepth--
			return result
		}
	} else {

		if a.OpOverload != nil {
			evalFnc = func(ctx *Context) bool {
				return b.OpOverload.StringNotEquals(ctx, ea)
			}
		} else {

			evalFnc = func(ctx *Context) bool {
				return ea != eb(ctx)
			}

		}

	}

	return &BoolEvaluator{
		EvalFnc:   evalFnc,
		isPartial: isPartialLeaf,
	}, nil
}

func BoolEquals(a *BoolEvaluator, b *BoolEvaluator, opts *Opts, state *state) (*BoolEvaluator, error) {
	partialA, partialB := a.isPartial, b.isPartial

	if a.EvalFnc == nil || (a.Field != "" && a.Field != state.field) {
		partialA = true
	}
	if b.EvalFnc == nil || (b.Field != "" && b.Field != state.field) {
		partialB = true
	}
	isPartialLeaf := partialA && partialB

	if a.Field != "" && b.Field != "" {
		isPartialLeaf = true
	}

	if (a.EvalFnc != nil || a.OpOverload != nil) && (b.EvalFnc != nil || b.OpOverload != nil) {
		ea, eb := a.EvalFnc, b.EvalFnc

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb(ctx)
				result := op1 == op2
				ctx.Logf("Evaluating %v == %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return a.OpOverload.BoolEquals(ctx, eb(ctx))
				}
			} else if b.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return b.OpOverload.BoolEquals(ctx, ea(ctx))
				}
			} else {

				evalFnc = func(ctx *Context) bool {
					return ea(ctx) == eb(ctx)
				}

			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc == nil && a.OpOverload == nil && b.EvalFnc == nil && b.OpOverload == nil {
		ea, eb := a.Value, b.Value

		return &BoolEvaluator{
			Value:     ea == eb,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc != nil || a.OpOverload != nil {
		ea, eb := a.EvalFnc, b.Value

		if a.Field != "" {
			if err := state.UpdateFieldValues(a.Field, FieldValue{Value: eb, Type: ScalarValueType}); err != nil {
				return nil, err
			}
		}

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb
				result := op1 == op2
				ctx.Logf("Evaluating %v == %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return a.OpOverload.BoolEquals(ctx, eb)
				}
			} else {

				evalFnc = func(ctx *Context) bool {
					return ea(ctx) == eb
				}

			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	ea, eb := a.Value, b.EvalFnc

	if b.Field != "" {
		if err := state.UpdateFieldValues(b.Field, FieldValue{Value: ea, Type: ScalarValueType}); err != nil {
			return nil, err
		}
	}

	var evalFnc func(ctx *Context) bool
	if opts.Debug {
		evalFnc = func(ctx *Context) bool {
			ctx.evalDepth++
			op1, op2 := ea, eb(ctx)
			result := op1 == op2
			ctx.Logf("Evaluating %v == %v => %v", op1, op2, result)
			ctx.evalDepth--
			return result
		}
	} else {

		if a.OpOverload != nil {
			evalFnc = func(ctx *Context) bool {
				return b.OpOverload.BoolEquals(ctx, ea)
			}
		} else {

			evalFnc = func(ctx *Context) bool {
				return ea == eb(ctx)
			}

		}

	}

	return &BoolEvaluator{
		EvalFnc:   evalFnc,
		isPartial: isPartialLeaf,
	}, nil
}

func BoolNotEquals(a *BoolEvaluator, b *BoolEvaluator, opts *Opts, state *state) (*BoolEvaluator, error) {
	partialA, partialB := a.isPartial, b.isPartial

	if a.EvalFnc == nil || (a.Field != "" && a.Field != state.field) {
		partialA = true
	}
	if b.EvalFnc == nil || (b.Field != "" && b.Field != state.field) {
		partialB = true
	}
	isPartialLeaf := partialA && partialB

	if a.Field != "" && b.Field != "" {
		isPartialLeaf = true
	}

	if (a.EvalFnc != nil || a.OpOverload != nil) && (b.EvalFnc != nil || b.OpOverload != nil) {
		ea, eb := a.EvalFnc, b.EvalFnc

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb(ctx)
				result := op1 != op2
				ctx.Logf("Evaluating %v != %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return a.OpOverload.BoolNotEquals(ctx, eb(ctx))
				}
			} else if b.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return b.OpOverload.BoolNotEquals(ctx, ea(ctx))
				}
			} else {

				evalFnc = func(ctx *Context) bool {
					return ea(ctx) != eb(ctx)
				}

			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc == nil && a.OpOverload == nil && b.EvalFnc == nil && b.OpOverload == nil {
		ea, eb := a.Value, b.Value

		return &BoolEvaluator{
			Value:     ea != eb,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc != nil || a.OpOverload != nil {
		ea, eb := a.EvalFnc, b.Value

		if a.Field != "" {
			if err := state.UpdateFieldValues(a.Field, FieldValue{Value: eb, Type: ScalarValueType}); err != nil {
				return nil, err
			}
		}

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb
				result := op1 != op2
				ctx.Logf("Evaluating %v != %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return a.OpOverload.BoolNotEquals(ctx, eb)
				}
			} else {

				evalFnc = func(ctx *Context) bool {
					return ea(ctx) != eb
				}

			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	ea, eb := a.Value, b.EvalFnc

	if b.Field != "" {
		if err := state.UpdateFieldValues(b.Field, FieldValue{Value: ea, Type: ScalarValueType}); err != nil {
			return nil, err
		}
	}

	var evalFnc func(ctx *Context) bool
	if opts.Debug {
		evalFnc = func(ctx *Context) bool {
			ctx.evalDepth++
			op1, op2 := ea, eb(ctx)
			result := op1 != op2
			ctx.Logf("Evaluating %v != %v => %v", op1, op2, result)
			ctx.evalDepth--
			return result
		}
	} else {

		if a.OpOverload != nil {
			evalFnc = func(ctx *Context) bool {
				return b.OpOverload.BoolNotEquals(ctx, ea)
			}
		} else {

			evalFnc = func(ctx *Context) bool {
				return ea != eb(ctx)
			}

		}

	}

	return &BoolEvaluator{
		EvalFnc:   evalFnc,
		isPartial: isPartialLeaf,
	}, nil
}

func GreaterThan(a *IntEvaluator, b *IntEvaluator, opts *Opts, state *state) (*BoolEvaluator, error) {
	partialA, partialB := a.isPartial, b.isPartial

	if a.EvalFnc == nil || (a.Field != "" && a.Field != state.field) {
		partialA = true
	}
	if b.EvalFnc == nil || (b.Field != "" && b.Field != state.field) {
		partialB = true
	}
	isPartialLeaf := partialA && partialB

	if a.Field != "" && b.Field != "" {
		isPartialLeaf = true
	}

	if (a.EvalFnc != nil || a.OpOverload != nil) && (b.EvalFnc != nil || b.OpOverload != nil) {
		ea, eb := a.EvalFnc, b.EvalFnc

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb(ctx)
				result := op1 > op2
				ctx.Logf("Evaluating %v > %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return a.OpOverload.GreaterThan(ctx, eb(ctx))
				}
			} else if b.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return b.OpOverload.GreaterThan(ctx, ea(ctx))
				}
			} else {

				evalFnc = func(ctx *Context) bool {
					return ea(ctx) > eb(ctx)
				}

			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc == nil && a.OpOverload == nil && b.EvalFnc == nil && b.OpOverload == nil {
		ea, eb := a.Value, b.Value

		return &BoolEvaluator{
			Value:     ea > eb,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc != nil || a.OpOverload != nil {
		ea, eb := a.EvalFnc, b.Value

		if a.Field != "" {
			if err := state.UpdateFieldValues(a.Field, FieldValue{Value: eb, Type: ScalarValueType}); err != nil {
				return nil, err
			}
		}

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb
				result := op1 > op2
				ctx.Logf("Evaluating %v > %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return a.OpOverload.GreaterThan(ctx, eb)
				}
			} else {

				evalFnc = func(ctx *Context) bool {
					return ea(ctx) > eb
				}

			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	ea, eb := a.Value, b.EvalFnc

	if b.Field != "" {
		if err := state.UpdateFieldValues(b.Field, FieldValue{Value: ea, Type: ScalarValueType}); err != nil {
			return nil, err
		}
	}

	var evalFnc func(ctx *Context) bool
	if opts.Debug {
		evalFnc = func(ctx *Context) bool {
			ctx.evalDepth++
			op1, op2 := ea, eb(ctx)
			result := op1 > op2
			ctx.Logf("Evaluating %v > %v => %v", op1, op2, result)
			ctx.evalDepth--
			return result
		}
	} else {

		if a.OpOverload != nil {
			evalFnc = func(ctx *Context) bool {
				return b.OpOverload.GreaterThan(ctx, ea)
			}
		} else {

			evalFnc = func(ctx *Context) bool {
				return ea > eb(ctx)
			}

		}

	}

	return &BoolEvaluator{
		EvalFnc:   evalFnc,
		isPartial: isPartialLeaf,
	}, nil
}

func GreaterOrEqualThan(a *IntEvaluator, b *IntEvaluator, opts *Opts, state *state) (*BoolEvaluator, error) {
	partialA, partialB := a.isPartial, b.isPartial

	if a.EvalFnc == nil || (a.Field != "" && a.Field != state.field) {
		partialA = true
	}
	if b.EvalFnc == nil || (b.Field != "" && b.Field != state.field) {
		partialB = true
	}
	isPartialLeaf := partialA && partialB

	if a.Field != "" && b.Field != "" {
		isPartialLeaf = true
	}

	if (a.EvalFnc != nil || a.OpOverload != nil) && (b.EvalFnc != nil || b.OpOverload != nil) {
		ea, eb := a.EvalFnc, b.EvalFnc

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb(ctx)
				result := op1 >= op2
				ctx.Logf("Evaluating %v >= %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return a.OpOverload.GreaterOrEqualThan(ctx, eb(ctx))
				}
			} else if b.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return b.OpOverload.GreaterOrEqualThan(ctx, ea(ctx))
				}
			} else {

				evalFnc = func(ctx *Context) bool {
					return ea(ctx) >= eb(ctx)
				}

			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc == nil && a.OpOverload == nil && b.EvalFnc == nil && b.OpOverload == nil {
		ea, eb := a.Value, b.Value

		return &BoolEvaluator{
			Value:     ea >= eb,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc != nil || a.OpOverload != nil {
		ea, eb := a.EvalFnc, b.Value

		if a.Field != "" {
			if err := state.UpdateFieldValues(a.Field, FieldValue{Value: eb, Type: ScalarValueType}); err != nil {
				return nil, err
			}
		}

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb
				result := op1 >= op2
				ctx.Logf("Evaluating %v >= %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return a.OpOverload.GreaterOrEqualThan(ctx, eb)
				}
			} else {

				evalFnc = func(ctx *Context) bool {
					return ea(ctx) >= eb
				}

			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	ea, eb := a.Value, b.EvalFnc

	if b.Field != "" {
		if err := state.UpdateFieldValues(b.Field, FieldValue{Value: ea, Type: ScalarValueType}); err != nil {
			return nil, err
		}
	}

	var evalFnc func(ctx *Context) bool
	if opts.Debug {
		evalFnc = func(ctx *Context) bool {
			ctx.evalDepth++
			op1, op2 := ea, eb(ctx)
			result := op1 >= op2
			ctx.Logf("Evaluating %v >= %v => %v", op1, op2, result)
			ctx.evalDepth--
			return result
		}
	} else {

		if a.OpOverload != nil {
			evalFnc = func(ctx *Context) bool {
				return b.OpOverload.GreaterOrEqualThan(ctx, ea)
			}
		} else {

			evalFnc = func(ctx *Context) bool {
				return ea >= eb(ctx)
			}

		}

	}

	return &BoolEvaluator{
		EvalFnc:   evalFnc,
		isPartial: isPartialLeaf,
	}, nil
}

func LesserThan(a *IntEvaluator, b *IntEvaluator, opts *Opts, state *state) (*BoolEvaluator, error) {
	partialA, partialB := a.isPartial, b.isPartial

	if a.EvalFnc == nil || (a.Field != "" && a.Field != state.field) {
		partialA = true
	}
	if b.EvalFnc == nil || (b.Field != "" && b.Field != state.field) {
		partialB = true
	}
	isPartialLeaf := partialA && partialB

	if a.Field != "" && b.Field != "" {
		isPartialLeaf = true
	}

	if (a.EvalFnc != nil || a.OpOverload != nil) && (b.EvalFnc != nil || b.OpOverload != nil) {
		ea, eb := a.EvalFnc, b.EvalFnc

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb(ctx)
				result := op1 < op2
				ctx.Logf("Evaluating %v < %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return a.OpOverload.LesserThan(ctx, eb(ctx))
				}
			} else if b.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return b.OpOverload.LesserThan(ctx, ea(ctx))
				}
			} else {

				evalFnc = func(ctx *Context) bool {
					return ea(ctx) < eb(ctx)
				}

			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc == nil && a.OpOverload == nil && b.EvalFnc == nil && b.OpOverload == nil {
		ea, eb := a.Value, b.Value

		return &BoolEvaluator{
			Value:     ea < eb,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc != nil || a.OpOverload != nil {
		ea, eb := a.EvalFnc, b.Value

		if a.Field != "" {
			if err := state.UpdateFieldValues(a.Field, FieldValue{Value: eb, Type: ScalarValueType}); err != nil {
				return nil, err
			}
		}

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb
				result := op1 < op2
				ctx.Logf("Evaluating %v < %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return a.OpOverload.LesserThan(ctx, eb)
				}
			} else {

				evalFnc = func(ctx *Context) bool {
					return ea(ctx) < eb
				}

			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	ea, eb := a.Value, b.EvalFnc

	if b.Field != "" {
		if err := state.UpdateFieldValues(b.Field, FieldValue{Value: ea, Type: ScalarValueType}); err != nil {
			return nil, err
		}
	}

	var evalFnc func(ctx *Context) bool
	if opts.Debug {
		evalFnc = func(ctx *Context) bool {
			ctx.evalDepth++
			op1, op2 := ea, eb(ctx)
			result := op1 < op2
			ctx.Logf("Evaluating %v < %v => %v", op1, op2, result)
			ctx.evalDepth--
			return result
		}
	} else {

		if a.OpOverload != nil {
			evalFnc = func(ctx *Context) bool {
				return b.OpOverload.LesserThan(ctx, ea)
			}
		} else {

			evalFnc = func(ctx *Context) bool {
				return ea < eb(ctx)
			}

		}

	}

	return &BoolEvaluator{
		EvalFnc:   evalFnc,
		isPartial: isPartialLeaf,
	}, nil
}

func LesserOrEqualThan(a *IntEvaluator, b *IntEvaluator, opts *Opts, state *state) (*BoolEvaluator, error) {
	partialA, partialB := a.isPartial, b.isPartial

	if a.EvalFnc == nil || (a.Field != "" && a.Field != state.field) {
		partialA = true
	}
	if b.EvalFnc == nil || (b.Field != "" && b.Field != state.field) {
		partialB = true
	}
	isPartialLeaf := partialA && partialB

	if a.Field != "" && b.Field != "" {
		isPartialLeaf = true
	}

	if (a.EvalFnc != nil || a.OpOverload != nil) && (b.EvalFnc != nil || b.OpOverload != nil) {
		ea, eb := a.EvalFnc, b.EvalFnc

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb(ctx)
				result := op1 <= op2
				ctx.Logf("Evaluating %v <= %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return a.OpOverload.LesserOrEqualThan(ctx, eb(ctx))
				}
			} else if b.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return b.OpOverload.LesserOrEqualThan(ctx, ea(ctx))
				}
			} else {

				evalFnc = func(ctx *Context) bool {
					return ea(ctx) <= eb(ctx)
				}

			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc == nil && a.OpOverload == nil && b.EvalFnc == nil && b.OpOverload == nil {
		ea, eb := a.Value, b.Value

		return &BoolEvaluator{
			Value:     ea <= eb,
			isPartial: isPartialLeaf,
		}, nil
	}

	if a.EvalFnc != nil || a.OpOverload != nil {
		ea, eb := a.EvalFnc, b.Value

		if a.Field != "" {
			if err := state.UpdateFieldValues(a.Field, FieldValue{Value: eb, Type: ScalarValueType}); err != nil {
				return nil, err
			}
		}

		var evalFnc func(ctx *Context) bool
		if opts.Debug {
			evalFnc = func(ctx *Context) bool {
				ctx.evalDepth++
				op1, op2 := ea(ctx), eb
				result := op1 <= op2
				ctx.Logf("Evaluating %v <= %v => %v", op1, op2, result)
				ctx.evalDepth--
				return result
			}
		} else {

			if a.OpOverload != nil {
				evalFnc = func(ctx *Context) bool {
					return a.OpOverload.LesserOrEqualThan(ctx, eb)
				}
			} else {

				evalFnc = func(ctx *Context) bool {
					return ea(ctx) <= eb
				}

			}

		}

		return &BoolEvaluator{
			EvalFnc:   evalFnc,
			isPartial: isPartialLeaf,
		}, nil
	}

	ea, eb := a.Value, b.EvalFnc

	if b.Field != "" {
		if err := state.UpdateFieldValues(b.Field, FieldValue{Value: ea, Type: ScalarValueType}); err != nil {
			return nil, err
		}
	}

	var evalFnc func(ctx *Context) bool
	if opts.Debug {
		evalFnc = func(ctx *Context) bool {
			ctx.evalDepth++
			op1, op2 := ea, eb(ctx)
			result := op1 <= op2
			ctx.Logf("Evaluating %v <= %v => %v", op1, op2, result)
			ctx.evalDepth--
			return result
		}
	} else {

		if a.OpOverload != nil {
			evalFnc = func(ctx *Context) bool {
				return b.OpOverload.LesserOrEqualThan(ctx, ea)
			}
		} else {

			evalFnc = func(ctx *Context) bool {
				return ea <= eb(ctx)
			}

		}

	}

	return &BoolEvaluator{
		EvalFnc:   evalFnc,
		isPartial: isPartialLeaf,
	}, nil
}
