package golyphone

import "testing"

func LatokirTest(t testing.T){
	s := latokir("Привет Мир")

	if s != "ПРИВЕТ МИР" {
		t.Error("Expected ПРИВЕТ МИР, got ", s)
	}

	s = latokir("Я так устал")

	if s != "Я ТАК УСТАЛ" {
		t.Error("Expected Я ТАК УСТАЛ, got ", s)
	}

	s = latokir("ХOЧУ")

	if s != "ХОЧУ" {
		t.Error("Expected ХОЧУ, got ", s)
	}

	s = latokir("Гyлять")

	if s != "ГУЛЯТЬ" {
		t.Error("Expected ГУЛЯТЬ, got ", s)
	}
}
