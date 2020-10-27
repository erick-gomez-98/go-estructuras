package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type contenidoWeb struct {
	multimedia []multimedia
}

func (cw *contenidoWeb) mostrar() string {
	var outStr string
	for _, c := range cw.multimedia {
		outStr += c.mostrar()
	}
	return outStr
}

type multimedia interface {
	mostrar() string
}

type imagen struct {
	titulo  string
	formato string
	canales string
}

func (i imagen) mostrar() string {
	return "imagen(titulo=" + i.titulo + ", formato=" + i.formato + ", canales=" + i.canales + ")\n"
}

type audio struct {
	titulo   string
	formato  string
	duracion string
}

func (a audio) mostrar() string {
	return "audio(titulo=" + a.titulo + ", formato=" + a.formato + ", duracion=" + a.duracion + ")\n"
}

type video struct {
	titulo  string
	formato string
	frames  string
}

func (v video) mostrar() string {
	return "video(titulo=" + v.titulo + ", formato=" + v.formato + ", frames=" + v.frames + ")\n"
}

func main() {
	cw := contenidoWeb{}

	var option int
	out := false
	inputReader := bufio.NewReader(os.Stdin)

	for out == false {
		fmt.Println("Selecciona una opcion")
		fmt.Println("1. Agregar imagen")
		fmt.Println("2. Agregar audio")
		fmt.Println("3. Agregar video")
		fmt.Println("4. Mostrar")
		fmt.Println("5. Salir")
		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Println("Ingresa el titulo de la imagen:")
			t, _ := inputReader.ReadString('\n')
			t = strings.TrimRight(t, "\r\n")

			fmt.Println("Ingresa el formato de la imagen:")
			f, _ := inputReader.ReadString('\n')
			f = strings.TrimRight(f, "\r\n")

			fmt.Println("Ingresa los canales de la imagen:")
			c, _ := inputReader.ReadString('\n')
			c = strings.TrimRight(c, "\r\n")

			cw.multimedia = append(cw.multimedia, imagen{t, f, c})
		case 2:
			fmt.Println("Ingresa el titulo del audio:")
			t, _ := inputReader.ReadString('\n')
			t = strings.TrimRight(t, "\r\n")

			fmt.Println("Ingresa el formato del audio:")
			f, _ := inputReader.ReadString('\n')
			f = strings.TrimRight(f, "\r\n")

			fmt.Println("Ingresa la duracion del audio:")
			c, _ := inputReader.ReadString('\n')
			c = strings.TrimRight(c, "\r\n")

			cw.multimedia = append(cw.multimedia, audio{t, f, c})
		case 3:
			fmt.Println("Ingresa el titulo del video:")
			t, _ := inputReader.ReadString('\n')
			t = strings.TrimRight(t, "\r\n")

			fmt.Println("Ingresa el formato del video:")
			f, _ := inputReader.ReadString('\n')
			f = strings.TrimRight(f, "\r\n")

			fmt.Println("Ingresa los frames del video:")
			c, _ := inputReader.ReadString('\n')
			c = strings.TrimRight(c, "\r\n")

			cw.multimedia = append(cw.multimedia, video{t, f, c})
		case 4:
			fmt.Println(cw.mostrar())
		case 5:
			out = true
		default:
			fmt.Println("Selecciona una opcion correcta")

		}
	}
}
