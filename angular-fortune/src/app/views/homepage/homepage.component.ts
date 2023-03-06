import { Component, OnInit } from '@angular/core';
import Pusher from 'pusher-js';
import * as PusherTypes from 'pusher-js';

import {HttpClient} from "@angular/common/http";

@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.css']
})
export class HomepageComponent implements OnInit{
  username = 'username';
  message = '';
  messages = [];

  constructor(private http: HttpClient) {}

  ngOnInit(): void{
    // Enable pusher logging - don't include this in production
    Pusher.logToConsole = true;

    const pusher = new Pusher('a621a1a5218dda4b051a', {
      cluster: 'us2'
    });

    const channel = pusher.subscribe('chat');
    channel.bind('message', (data: any) => {
      this.messages.push(data);
    });
  }

  submit(): void {
    this.http.post('http://localhost:8000/api/messages', {
      username: this.username,
      message: this.message
    }).subscribe(() => this.message = '');
  }
}
