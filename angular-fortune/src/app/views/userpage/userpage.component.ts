import { Component } from '@angular/core';
import { SocialAuthService } from "@abacritt/angularx-social-login";

import Pusher from 'pusher-js';

import {HttpClient} from "@angular/common/http";

@Component({
  selector: 'app-userpage',
  templateUrl: './userpage.component.html',
  styles: [
  ]
})

export class UserpageComponent{

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

    const pusher = new Pusher('a621a1a5218dda4b051a', {
      cluster: 'us2'
    });

    //updating values only if the user is logged in.
    if (this.loggedIn){
      this.http.post('http://localhost:8000/api/user/populate', {
        //When submit is called, it will sent this usename and message to the backend. 
        username: this.user.name,
        email: this.user.email,
        userid: this.user.id
      }).subscribe();
    }

  }

  getData(): void {
    //alert(JSON.stringify("This is working"));
    //will get the variables from backend. res is the response

    //the get request below is for receiving the last fortune.
    // //!Eventually, replace this with all the fortunes. Later we can have a get request for updating this from the backend and it'll be easier
    // this.http.get('http://localhost:8000/api/user/frontend/fid').subscribe(
    //   (data : any) => {
    //     //if the array is empty, add the first item
    //     if (!this.fid.length){
    //       this.fid.push(data);
    //     }
    //     //if its not empty, it checks if the last item is the same
    //     else if (data != this.fid[this.fid.length - 1]){
    //       this.fid.push(data);
    //     }
    //   }
    // );
    // this.http.get('http://localhost:8000/api/user/frontend/fid').subscribe(
    //   (data : any) => {
    //     this.fid.push(data);
    //   }
    // );

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