module Main exposing (..)

import Html exposing (..)
import Browser exposing (sandbox)
import Html.Events exposing (..)


-- type alias Product =
--     { name : String
--     , model : String
--     , cost : Int
--     }
-- products =
--     [ { name = "ValueName1", model = "", cost = 11 }
--     , { name = "ValueName2", model = "", cost = 2 }
--     , { name = "ValueName3", model = "", cost = 9000 }
--     ]
-- renderProduct : Product -> Html msg
-- renderProduct product =
--     li []
--         [ text product.name
--         , text ", "
--         , b []
--             [ text <| String.fromInt product.cost ]
--         ]
-- renderProducts : List Product -> Html msg
-- renderProducts prods =
--     div
--         [ style "font-family" "-apple-system"
--         , style "padding" "1em"
--         ]
--         [ h1 [] [ text "Products" ]
--         , ul [] (List.map renderProduct prods)
--         ]


model =
    { showFace = False }


type Msg
    = ShowFace


update msg model_ =
    case msg of
        ShowFace ->
            { model_ | showFace = True }


view model_ =
    div []
        [ h1 [] [ text " FACE GENERATOR 3000 " ]
        , button [ onClick ShowFace ] [ text "Generate Face" ]
        , if model_.showFace then
            text "`~`"
          else
            text ""
        ]


main =
    Browser.sandbox
        { init = model
        , update = update
        , view = view
        }
