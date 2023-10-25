package api


type ApiServer interface {

}


type JsonApiServer struct {
  port int
}

type GrpcServer struct {

}

func NewServer(port int) (*JsonApiServer) {
  return &JsonApiServer{
    port : port
  }
}
