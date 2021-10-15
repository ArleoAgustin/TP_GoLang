# TP_GoLang


Para cumplir con la consigna solicitada pense lo siguiente:

archivo main

{
    Dentro de la funcion main se crea una variable para almacenar la cadena,
    esta luego se pasa por parametro en la funcion SetParametro. Set parametro
    retorna dos variables una que va a ser el objeto y la otra es por si se
    genera un error. Si el error es nil, es decir que no hubo error en la
    cadena se imprime por pantalla el objeto. En el caso contrario se imprime
    un mensaje de error
}

archivo entities

{
    Es el encargado de crear la estructura del objeto y realizar todo el procedimiento
    de verificacion.
    La funcion NewResult es la encargada de crear el objeto y retornarlo.
    SetCadena se encarga de obtener los datos de la cadena (Type, Length, Value),
    luego se llama a NewResult para asi crear el objeto. A continuacion verifica
    si el Value coincide con el Length, si esto sucede entra en accion la funcion esCorrecto. En
    el caso de que esCorrecto devuelva True se retornara el objeto junto con un nil, ya
    que todo salio bien. En caso de que haya un error retorna un objeto vacio,
    junto con un error.
    La funcion esCorrecto es la encargada de verificar si la cadena es de tipo alfabetica
    o numerica. En ambos casos se accede a el Value del objeto, para asi recorrer esa porcion de la cadena. Dependiendo el tipo de cadena va a verificar si cada indice es de tipo digito o
    si es una letra, si llega a haber un digito mezclado entre las letras este va
    a cortar y dar error. Lo mismo en en caso contrario, el value puede
    ser solo de un tipo digito o letra. Si llego al final de esa porcion de cadena va a retornar true, si esto
    no es asi retorna false.
}

Asi es como lo pense para poder cumplir con la consigna.
Al realizar el test dio un coverage del 100%
    
                                                   By Agustin Arleo