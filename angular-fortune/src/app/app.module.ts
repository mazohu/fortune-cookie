import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { SocialLoginModule, SocialAuthServiceConfig } from '@abacritt/angularx-social-login';
import {
  GoogleLoginProvider,
} from '@abacritt/angularx-social-login';

//routingComponents will automatically import all the components we rout to. All we need to do is update this list
import { AppRoutingModule, routingComponents } from './app-routing.module';
import { AppComponent } from './app.component';

import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { EatcookieComponent } from './views/eatcookie/eatcookie.component';
import { PastfortunesComponent } from './views/pastfortunes/pastfortunes.component';
import { UserprofileComponent } from './views/userprofile/userprofile.component';

@NgModule({
  declarations: [
    AppComponent,
    routingComponents,
    EatcookieComponent,
    PastfortunesComponent,
    UserprofileComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    SocialLoginModule,
    FormsModule,
    HttpClientModule,
  ],
  providers: [
    {
      provide: 'SocialAuthServiceConfig',
      useValue: {
        autoLogin: false,
        providers: [
          {
            id: GoogleLoginProvider.PROVIDER_ID,
            provider: new GoogleLoginProvider(
              '152384338353-hbdhivcu978k31fsq9d9fnsjvgaktkk1.apps.googleusercontent.com'
            )
          },
        ],
        onError: (err) => {
          console.error(err);
        }
      } as SocialAuthServiceConfig,
    }
  ],
  
  bootstrap: [AppComponent]
})
export class AppModule { }
