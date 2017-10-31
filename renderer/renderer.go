package renderer

//Order in which renderables are rendered
type Order int

//RenderableWorld represnets an object with a world pos
type RenderableWorld interface {
}

//RenderableScreen represents a gui item
type RenderableScreen interface {
}

//Renderer is responsible for rendering everything in a scene
type Renderer struct {
}
