/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import "github.com/91diego/platform-challenge/cmd"

// longuitud de la calle envio destino es PAR;
// el puntaje de idoniedad (SS) = numero de VOCALES del nombre del conductor * 1.5

// longuitud de la calle envio destino es IMPAR;
// el puntaje de idoniedad (SS) = numero de CONSONANTES del nombre del conductor * 1

// longuitud de la calle envio destino tiene algo en COMUN con el nombre del conductor;
// el puntaje de idoniedad (SS) = alguno de los casos anteriores + 50%
// ejemplos;
// - misma longuitud entre nombre de la calle y conductos
// - nombre del conductor en la calle

func main() {
	cmd.Execute()
}
