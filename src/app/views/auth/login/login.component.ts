import { Component, OnInit } from "@angular/core";
import { SocialAuthService } from "angularx-social-login";

@Component({
  selector: "app-login",
  templateUrl: "./login.component.html",
})
export class LoginComponent implements OnInit {
  user: any;
  loggedIn: any;
  constructor(private authService: SocialAuthService) { }

  ngOnInit() {
    this.authService.authState.subscribe((user) => {
      this.user = user;
      this.loggedIn = (user != null);
    });
  }
}
