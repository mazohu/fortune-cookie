import { Component, OnInit } from '@angular/core';
import { OktaAuthStateService } from '@okta/okta-angular';
import { OktaAuth } from '@okta/okta-auth-js';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})

export class HomeComponent {
  title = 'home';

  constructor(private oktaAuth: OktaAuth, public authService: OktaAuthStateService) {
  }
  
  async login() {
    await this.oktaAuth.signInWithRedirect();
  }

  async logout() {
    await this.oktaAuth.signOut();
  }

}