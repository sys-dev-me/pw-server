package main

import "math/rand"



func  ( this *Application ) Init () {
	this.PWDB = make([]Password, 0)
}

func ( this *Application) SetVoc ( voc string ) {

	this.Set = voc

}

func ( this *Application ) Generate ( pwSize int ) (string) {

	pw := new(Password)
	pw.Len = pwSize

        src := []rune( this.Set )
        res := make([]rune, pwSize)

        for i := range res {
        	res[ i ] = src[ rand.Intn ( len ( src ) - 1 ) ]
    	}

	pw.Key = string( res )
	this.PWDB = append ( this.PWDB, *pw )


        return pw.Key
}

