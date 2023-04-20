import { Component } from '@angular/core';
import { SocialAuthService } from "@abacritt/angularx-social-login";
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
})
export class LoginComponent {
  title = 'angular-fortune';
  user:any;
  loggedIn:any;

  constructor(private authService: SocialAuthService, private router: Router){}

  ngOnInit(){
    this.authService.authState.subscribe((user) => {
      this.user = user;
      this.loggedIn = (user != null);
      console.log(this.user)
    });
  }

  navigate(){
    this.router.navigate(['/userpage'])
  }
}
