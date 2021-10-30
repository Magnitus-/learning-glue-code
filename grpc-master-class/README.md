# About

Exercises implementation for this course: https://www.udemy.com/course/grpc-golang

I'm pretty close to what the instructor is doing, though the **option go_package** line of the instructor in the proto definition didn't work for me so I had to fix that and make it the proper relative path.

I'm also using golang modules and using **replace** statements in the module definition to make client & server binaries work with a shared protobuf definition dependencie in a monorepo style.