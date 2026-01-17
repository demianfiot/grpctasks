testing by client or by grpcurl 

example: 

grpcurl -plaintext \
  -d '{"name":"Demian"}' \
  localhost:50051 user.UserService/CreateUser

grpcurl -plaintext localhost:50051 user.UserService/ListUsers

grpcurl -plaintext \
  -d '{"title":"Clean room","difficulty":"hard","xpReward":10}' \
  localhost:50051 task.TaskService/CreateTask
