import { Component, OnInit } from "@angular/core";
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: "app-login",
  templateUrl: "./login.component.html",
})
export class LoginComponent {
  constructor(
    private httpClient: HttpClient,
    private router: Router
  ) {}

  googleLogin(){
    this.httpClient.post("${environment.gateway}/auth/google", {
    }).subscribe((response: any) => {
      if(response){
        this.router.navigate(['../userhome'])
      }
    })
  }

  ngOnInit(): void {}
}
