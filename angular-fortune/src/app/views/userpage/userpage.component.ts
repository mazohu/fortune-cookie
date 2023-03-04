import { Component } from '@angular/core';
import { SocialAuthService } from "@abacritt/angularx-social-login";

@Component({
  selector: 'app-userpage',
  templateUrl: './userpage.component.html',
  styles: [
  ]
})
export class UserpageComponent {

  user:any;
  loggedIn:any;

  constructor(private authService: SocialAuthService){}

  ngOnInit(){
    this.authService.authState.subscribe((user) => {
      this.user = user;
      this.loggedIn = (user != null);
      console.log(this.user)
    });
  }

}
