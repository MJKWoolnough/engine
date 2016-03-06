package engine

type Key uint

const (
	KeyEscape       Key = iota //Escape
	KeyUp                      //Up
	KeyDown                    //Down
	KeyLeft                    //Left
	KeyRight                   //Right
	MouseLeft                  //Left
	MouseMiddle                //Middle
	MouseRight                 //Right
	MouseScrollUp              //Mouse Scroll Up
	MouseScrollDown            //Mouse Scroll Down
)
