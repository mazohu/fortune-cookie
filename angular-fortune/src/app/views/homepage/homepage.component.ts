import { Component, OnInit } from '@angular/core';
import Pusher from 'pusher-js';

import {HttpClient} from "@angular/common/http";

@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.css']
})
export class HomepageComponent implements OnInit{
  username : any = 'username';
  message : any = '';
  messages : any[] = [];

  constructor(private http: HttpClient) {}

  ngOnInit(): void{
    //Enable pusher logging - don't include this in production
    //Pusher.logToConsole = true;

    const pusher = new Pusher('a621a1a5218dda4b051a', {
      cluster: 'us2'
    });

    //the channel is chat
    const channel = pusher.subscribe('chat');

    //the event is 'message'
    channel.bind('message', (data : any) => {
      this.messages.push(data)
      //alert(JSON.stringify(data));
    });

    for(let i=0 ; i < this.messages.length ; i++){  //How to properly iterate here!!
      console.log(this.messages[i])
    }
  }

  submit(): void {
    this.http.post('http://localhost:8000/api/messages', {
      //When submit is called, it will sent this usename and message to the backend. 
      username: this.username,
      message: this.message
    }).subscribe();
  }

  changeFn(e : any) {
    this.message = e.target.value;
  }
}
