module Main exposing (..)

import Signal exposing (..)
import Graphics.Element exposing (show, Element)
import Mouse

position : Signal (Int, Float)
position = 
    Mouse.position

main: Signal Element
main =
    Signal.map show position

