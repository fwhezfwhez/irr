package irr

type WrapF struct {
	f func()

	ctx *Context
}

func WrapFunc(f func()) *WrapF {
	return &WrapF{
		f:   f,
		ctx: newContext(),
	}
}

func (wf *WrapF) Use(f func(c *Context)) {
	wf.ctx.addHandler(f)
}

func (wf *WrapF) Handle() {
	wf.ctx.handlers = append(wf.ctx.handlers, func(c *Context) {
		wf.f()
	})

	if len(wf.ctx.handlers) > 0 {
		wf.ctx.Next()
	}
	wf.ctx.Reset()
}
