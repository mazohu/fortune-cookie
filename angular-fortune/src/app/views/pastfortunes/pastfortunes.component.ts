import { Component } from '@angular/core';
import { SocialAuthService } from "@abacritt/angularx-social-login";

import Pusher from 'pusher-js';
import {HttpClient} from "@angular/common/http";

@Component({
  selector: 'app-pastfortunes',
  templateUrl: './pastfortunes.component.html',
  styleUrls: ['./pastfortunes.component.css']
})
export class PastfortunesComponent {
  user:any;
  loggedIn:any;

  //username, email, and id is all contained in user above
  fid : string[] = [];
  content : string = "";
  date : string = "";

  constructor(private authService: SocialAuthService, private http: HttpClient){}

  
  ngOnInit(){
    this.authService.authState.subscribe((user) => {
      this.user = user;
      this.loggedIn = (user != null);
      console.log(this.user)
    });

    //updating values only if the user is logged in.
    if (this.loggedIn){
      this.http.post('http://localhost:8000/api/user/populate', {
        //When submit is called, it will sent this usename and message to the backend. 
        username: this.user.name,
        email: this.user.email,
        userid: this.user.id
      }).subscribe(data => {
        this.receive();
      });
    }
  }

  receive():void{
    //receiving a fortune.
    //alert(JSON.stringify("This is working"));
    this.http.get('http://localhost:8000/api/user/frontend/fid').subscribe(
        (data : any) => {
          
          this.fid = data;
        }
    );
  }
}
