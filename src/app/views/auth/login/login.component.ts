import { Component, OnInit } from "@angular/core";
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { environment } from "src/environments/environment"
import { Injectable } from "@angular/core";
import {GoogleLoginProvider, SocialAuthService} from '@abacritt/angularx-social-login';


@Injectable({
  providedIn: "root",
})

@Component({
  selector: "app-login",
  templateUrl: "./login.component.html",
})
export class LoginComponent implements OnInit {
  constructor(
    private router: Router,
    private socialAuthService: SocialAuthService
  ) {}

  googleLogin(){
    this.socialAuthService.signIn(GoogleLoginProvider.PROVIDER_ID)
      .then(() => this.router.navigate(['mainpage']));
  }

  ngOnInit(): void {}
}
