module Api.Songservice exposing (..)

-- DO NOT EDIT
-- AUTOGENERATED BY THE ELM PROTOCOL BUFFER COMPILER
-- https://github.com/tiziano88/elm-protobuf
-- source file: api/songservice.proto

import Protobuf exposing (..)

import Json.Decode as JD
import Json.Encode as JE


uselessDeclarationToPreventErrorDueToEmptyOutputFile = 42


type alias AddSongRequest =
    { newSong : Maybe Song -- 1
    }


addSongRequestDecoder : JD.Decoder AddSongRequest
addSongRequestDecoder =
    JD.lazy <| \_ -> decode AddSongRequest
        |> optional "newSong" songDecoder


addSongRequestEncoder : AddSongRequest -> JE.Value
addSongRequestEncoder v =
    JE.object <| List.filterMap identity <|
        [ (optionalEncoder "newSong" songEncoder v.newSong)
        ]


type alias UpdateSongRequest =
    { songID : String -- 1
    , updatedSong : Maybe Song -- 2
    }


updateSongRequestDecoder : JD.Decoder UpdateSongRequest
updateSongRequestDecoder =
    JD.lazy <| \_ -> decode UpdateSongRequest
        |> required "songID" JD.string ""
        |> optional "updatedSong" songDecoder


updateSongRequestEncoder : UpdateSongRequest -> JE.Value
updateSongRequestEncoder v =
    JE.object <| List.filterMap identity <|
        [ (requiredFieldEncoder "songID" JE.string "" v.songID)
        , (optionalEncoder "updatedSong" songEncoder v.updatedSong)
        ]


type alias GetSongRequest =
    { name : String -- 1
    , artists : List String -- 2
    , songID : String -- 3
    }


getSongRequestDecoder : JD.Decoder GetSongRequest
getSongRequestDecoder =
    JD.lazy <| \_ -> decode GetSongRequest
        |> required "name" JD.string ""
        |> repeated "artists" JD.string
        |> required "songID" JD.string ""


getSongRequestEncoder : GetSongRequest -> JE.Value
getSongRequestEncoder v =
    JE.object <| List.filterMap identity <|
        [ (requiredFieldEncoder "name" JE.string "" v.name)
        , (repeatedFieldEncoder "artists" JE.string v.artists)
        , (requiredFieldEncoder "songID" JE.string "" v.songID)
        ]


type alias GetSongResponse =
    { result : List GetSongResponse_GetSongResult -- 1
    }


getSongResponseDecoder : JD.Decoder GetSongResponse
getSongResponseDecoder =
    JD.lazy <| \_ -> decode GetSongResponse
        |> repeated "result" getSongResponse_GetSongResultDecoder


getSongResponseEncoder : GetSongResponse -> JE.Value
getSongResponseEncoder v =
    JE.object <| List.filterMap identity <|
        [ (repeatedFieldEncoder "result" getSongResponse_GetSongResultEncoder v.result)
        ]


type alias GetSongResponse_GetSongResult =
    { song : Maybe Song -- 1
    , songID : String -- 2
    }


getSongResponse_GetSongResultDecoder : JD.Decoder GetSongResponse_GetSongResult
getSongResponse_GetSongResultDecoder =
    JD.lazy <| \_ -> decode GetSongResponse_GetSongResult
        |> optional "song" songDecoder
        |> required "songID" JD.string ""


getSongResponse_GetSongResultEncoder : GetSongResponse_GetSongResult -> JE.Value
getSongResponse_GetSongResultEncoder v =
    JE.object <| List.filterMap identity <|
        [ (optionalEncoder "song" songEncoder v.song)
        , (requiredFieldEncoder "songID" JE.string "" v.songID)
        ]


type alias DeleteSongRequest =
    { songID : String -- 1
    }


deleteSongRequestDecoder : JD.Decoder DeleteSongRequest
deleteSongRequestDecoder =
    JD.lazy <| \_ -> decode DeleteSongRequest
        |> required "SongID" JD.string ""


deleteSongRequestEncoder : DeleteSongRequest -> JE.Value
deleteSongRequestEncoder v =
    JE.object <| List.filterMap identity <|
        [ (requiredFieldEncoder "SongID" JE.string "" v.songID)
        ]


type alias Song =
    { name : String -- 1
    , artists : List String -- 2
    , audio : Bytes -- 3
    , audioformat : String -- 4
    }


songDecoder : JD.Decoder Song
songDecoder =
    JD.lazy <| \_ -> decode Song
        |> required "name" JD.string ""
        |> repeated "artists" JD.string
        |> required "audio" bytesFieldDecoder []
        |> required "audioformat" JD.string ""


songEncoder : Song -> JE.Value
songEncoder v =
    JE.object <| List.filterMap identity <|
        [ (requiredFieldEncoder "name" JE.string "" v.name)
        , (repeatedFieldEncoder "artists" JE.string v.artists)
        , (requiredFieldEncoder "audio" bytesFieldEncoder [] v.audio)
        , (requiredFieldEncoder "audioformat" JE.string "" v.audioformat)
        ]


type alias Empty =
    {
    }


emptyDecoder : JD.Decoder Empty
emptyDecoder =
    JD.lazy <| \_ -> decode Empty


emptyEncoder : Empty -> JE.Value
emptyEncoder v =
    JE.object <| List.filterMap identity <|
        [
        ]
