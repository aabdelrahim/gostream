import songservice_pb2_grpc
import songservice_pb2
import grpc
import logging



def delete_song(stub):
    return stub.Delete(songservice_pb2.DeleteSongRequest(SongID="ID-1234"))


def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel('localhost:8080') as channel:
        stub = songservice_pb2_grpc.SongServiceStub(channel)
        print("-------------- Delete Song --------------")
        delete_song(stub)


if __name__ == '__main__':
    logging.basicConfig()
    run()