import { Component, OnInit } from '@angular/core';

import OktaSignIn from '@okta/okta-signin-widget';
import oktaConfig from '../okta.config';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})

export class LoginComponent implements OnInit {
signIn: any;

constructor() {
  this.signIn = new OktaSignIn({
    baseUrl: oktaConfig.issuer.split('/oauth2')[0],
    clientId: oktaConfig.clientId,
    redirectUri: oktaConfig.redirectUri,
    i18n: {
      en: {
        'primaryauth.title': 'Sign in to Go Blog App',
      },
    },
    authParams: {
      responseType: ['id_token', 'token'],
      issuer: oktaConfig.issuer,
      display: 'page'
    },
  });
}

ngOnInit() {
  this.signIn.renderEl(
    { el: '#sign-in-widget' },
    () => {},
    (err: any) => {
      throw err;
    },
  );
}
}
