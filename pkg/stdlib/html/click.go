package html

import (
	"context"
	"github.com/MontFerret/ferret/pkg/runtime/core"
	"github.com/MontFerret/ferret/pkg/runtime/values"
	"github.com/MontFerret/ferret/pkg/stdlib/html/driver/dynamic"
)

/*
 * Dispatches click event on a given element
 * @param source (Document | Element) - Event source.
 * @param selector (String, optional) - Optional selector. Only used when a document instance is passed.
 */
func Click(_ context.Context, args ...core.Value) (core.Value, error) {
	err := core.ValidateArgs(args, 1, 2)

	if err != nil {
		return values.False, err
	}

	// CLICK(el)
	if len(args) == 1 {
		arg1 := args[0]

		err := core.ValidateType(arg1, core.HTMLElementType)

		if err != nil {
			return values.False, err
		}

		el, ok := arg1.(*dynamic.HTMLElement)

		if !ok {
			return values.False, core.Errors(core.ErrInvalidType, ErrNotDynamic)
		}

		return el.Click()
	}

	// CLICK(doc, selector)
	arg1 := args[0]
	selector := args[1].String()

	err = core.ValidateType(arg1, core.HTMLDocumentType)

	if err != nil {
		return values.None, err
	}

	doc, ok := arg1.(*dynamic.HTMLDocument)

	if !ok {
		return values.False, core.Errors(core.ErrInvalidType, ErrNotDynamic)
	}

	return doc.ClickBySelector(values.NewString(selector))
}
