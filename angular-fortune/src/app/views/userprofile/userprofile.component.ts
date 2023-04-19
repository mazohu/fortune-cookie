import { Component } from '@angular/core';
import { SocialAuthService } from '@abacritt/angularx-social-login';

import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-userprofile',
  templateUrl: './userprofile.component.html',
  styleUrls: ['./userprofile.component.css']
})

export class UserprofileComponent {

  //username, email, and ID are all contained in the user variable.
  user:any;
  loggedIn:any;

  submitted: boolean = false;
  lasttime: any = '';
  newFortune: string = '';

  constructor(private authService: SocialAuthService, private http: HttpClient){}

  ngOnInit() {
    this.authService.authState.subscribe((user) => {
      this.user = user;
      this.loggedIn = (user != null);
      console.log(this.user);
    });

    // updating values only if the user is logged in
    if (this.loggedIn){
      this.http.post('http://localhost:8000/api/user/populate', {
      
      // when submit is called, it will send these to the backend:
      username: this.user.name,
      email: this.user.email,
      userid: this.user.id
    }).subscribe(data => {
      this.getData();
    });
    }
  }

  getData(): void {
    this.http.get('http://localhost:8000/api/user/frontend/submitted').subscribe(
      (data : any) => {
        if (data == 1){
          this.submitted = true;
        }
        else{
          this.submitted = false;
        }
      }
    );

    this.http.get('http://localhost:8000/api/user/frontend/lastTime').subscribe(
      (data : any) => {
        this.lasttime = data;
      }
    );
  }
}