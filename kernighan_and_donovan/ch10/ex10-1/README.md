cd ../../ch3/mandelbrot
go build mandelbrot.go
cd ../../ch10/ex10-1
go build jpeg.go
../../ch3/mandelbrot/mandelbrot | ./jpeg --output jpeg >mandelbrot.jpg
../../ch3/mandelbrot/mandelbrot | ./jpeg --output gif >mandelbrot.gif
../../ch3/mandelbrot/mandelbrot | ./jpeg --output png >mandelbrot.png
