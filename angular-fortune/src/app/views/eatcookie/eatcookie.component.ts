import { Component } from '@angular/core';
import { SocialAuthService } from "@abacritt/angularx-social-login";

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
  submitted : boolean = false;
  lasttime : any = '';
  lastdate : string = '';
  newFortune : string = '';
  getFortune : string = '';

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

    this.http.get('http://localhost:8000/api/user/frontend/todayfortune').subscribe(
      (data : any) => {
        this.getFortune = data;
      }
    );

    this.http.get('http://localhost:8000/api/user/frontend/lastdate').subscribe(
      (data : any) => {
        this.lastdate = data;
      }
    );
  }

  receive():void{
    //receiving a fortune.
    //alert(JSON.stringify("This is working"));
    if (this.getFortune == ''){
      this.http.get('http://localhost:8000/api/user/frontend/getFortune').subscribe(
        (data : any) => {
          this.getFortune = data;
        }
      );
    }
    else{
      alert(JSON.stringify("Only one fortune a day dummy"));
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
        }).subscribe(data => {
          this.getData();
          this.newFortune = this.newFortune + ", it was submitted!" 
        });
    }
    else{
      alert(JSON.stringify("You can't get another fortune dummy"));
    }

  }

  changeFortune(e : any) {
    this.newFortune = e.target.value;
  }

}
