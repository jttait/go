cd ../../ch3/mandelbrot
go build mandelbrot.go
cd ../../ch10/jpeg
go build jpeg.go
../../ch3/mandelbrot/mandelbrot | ./jpeg >mandelbrot.jpg
