import { Component } from '@angular/core';
import { SocialAuthService } from "@abacritt/angularx-social-login";

import Pusher from 'pusher-js';
import {HttpClient} from "@angular/common/http";
@Component({
  selector: 'app-eatcookie',
  templateUrl: './eatcookie.component.html',
  styleUrls: ['./eatcookie.component.css']
})
export class EatcookieComponent {
  user:any;
  loggedIn:any;

  //username, email, and id is all contained in user above
  fid : string[] = [];
  submitted : boolean = false;
  lasttime : any = '';
  newFortune : string = '';

  constructor(private authService: SocialAuthService, private http: HttpClient){}

  ngOnInit(){
    this.authService.authState.subscribe((user) => {
      this.user = user;
      this.loggedIn = (user != null);
      console.log(this.user)
    });

    // const pusher = new Pusher('a621a1a5218dda4b051a', {
    //   cluster: 'us2'
    // });

    // //the channel is chat
    // const channel = pusher.subscribe('chat');

    // //the event is 'message'
    // channel.bind('message', (data : any) => {
    //   this.messages.push(data)
    //   //alert(JSON.stringify(data));
    // });

    // for(let i=0 ; i < this.messages.length ; i++){  //How to properly iterate here!!
    //   console.log(this.messages[i])
    // }

    // //updating values only if the user is logged in.
    if (this.loggedIn){
      this.http.post('http://localhost:8000/api/user/populate', {
        //When submit is called, it will sent this usename and message to the backend. 
        username: this.user.name,
        email: this.user.email,
        userid: this.user.id
      }).subscribe();
    }
  }

  submit(): void {
    if (!this.submitted){
      //when submitted is false, you're able to submit a fortune
      //updating values only if the user is logged in.
        this.http.post('http://localhost:8000/api/user/submitFortune', {
          //When submit is called, it will sent this usename and message to the backend. 
          //!Later find a way to input a new fortune and submit it here
          newfortune: this.newFortune
        }).subscribe();
        this.newFortune = "Our Fortune was Submitted"
    }
    else{
      alert(JSON.stringify("You can't get another fortune dummy"));
    }

  }

  changeFortune(e : any) {
    this.newFortune = e.target.value;
  }

}
