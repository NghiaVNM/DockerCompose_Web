package com.example.demo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class HelloController {

    // @Autowired
    // private PersonRepository personRepository;

    @GetMapping("/")
    public String hello() {
        return "Hello World for Java!";
    }

    // @GetMapping("/person")
    // public Person getPerson(@RequestParam String name) {
    //     return personRepository.findByName(name);
    // }
}