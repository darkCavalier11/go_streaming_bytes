# Goal
The main goal of the project is to receive a music id from the client along with a authentication key(encrypted from client with key). Upon receving the key and successful authentication, it will move to the next phase.
Upon receving the music id, it will fetch the url with code 140(i.e. m4a audio) and stream it to the client using grpc.(details are still need to refined).