import { Component } from '@angular/core';
import { SocialAuthService } from "@abacritt/angularx-social-login";
import { GoogleLoginProvider } from "@abacritt/angularx-social-login";
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'angular-fortune';

  // private accessToken = '';

  // constructor(private authService: SocialAuthService, private httpClient: HttpClient) { }

  // refreshToken(): void {
  //   this.authService.refreshAuthToken(GoogleLoginProvider.PROVIDER_ID);
  // }

  // getAccessToken(): void {
  //   this.authService.getAccessToken(GoogleLoginProvider.PROVIDER_ID).then(accessToken => this.accessToken = accessToken);
  // }

  // getGoogleCalendarData(): void {
  //   if (!this.accessToken) return;

  //   this.httpClient
  //     .get('https://www.googleapis.com/calendar/v3/calendars/primary/events', {
  //       headers: { Authorization: `Bearer ${this.accessToken}` },
  //     })
  //     .subscribe((events) => {
  //       alert('Look at your console');
  //       console.log('events', events);
  //     });
  // }

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
