import { Component, OnInit } from "@angular/core";
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { environment } from "src/environments/environment"
import { Injectable } from "@angular/core";
@Injectable({
  providedIn: "root",
})

@Component({
  selector: "app-login",
  templateUrl: "./login.component.html",
})
export class LoginComponent implements OnInit {
  constructor(
    private httpClient: HttpClient,
    private router: Router
  ) {}

  googleLogin(){
    return this.httpClient.get(environment.gateway + '/auth/google')
  }

  ngOnInit(): void {}
}
