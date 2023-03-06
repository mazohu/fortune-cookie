import { Component } from '@angular/core';
//import Pusher from 'pusher-js';
import {HttpClient} from "@angular/common/http";

@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.css']
})
export class HomepageComponent {
  username = 'username';
  message = '';
  messages = [];

  // constructor(private http: HttpClient) {}

  // ngOnInit(): void {
  //   Pusher.logToConsole = true;

  //   const pusher = new Pusher('25291c0752d6089a660c', {
  //     cluster: 'eu'
  //   });

  //   const channel = pusher.subscribe('chat');
  //   channel.bind('message', data => {
  //     this.messages.push(data);
  //   });
  // }

  submit(): void {
    // this.http.post('http://localhost:8000/api/messages', {
    //   username: this.username,
    //   message: this.message
    // }).subscribe(() => this.message = '');
  }
}
