import { Component, OnInit } from '@angular/core';
import { Location } from '@angular/common';
import { Router, ActivatedRoute, ParamMap } from '@angular/router';
import { NgForm } from '@angular/forms';
import { PostService } from 'src/app/shared/services/post.service';
import { IPost } from 'src/app/shared/models/post';
import { OktaAuthStateService } from '@okta/okta-angular';
import { OktaAuth } from '@okta/okta-auth-js';

import { switchMap } from 'rxjs/operators';
import { Subscription } from 'rxjs/internal/Subscription';

@Component({
  selector: 'app-edit-post',
  templateUrl: './edit-post.component.html',
  styleUrls: ['./edit-post.component.scss']
})
export class EditPostComponent implements OnInit {
  post: IPost;
  posts: IPost[];
  userEmail: string;
  regex: any = /new/g;

  constructor(
    private postService: PostService,
    public oktaAuth: OktaAuthStateService,
    private router: Router,
    private route: ActivatedRoute,
    private location: Location
  ) {
  }

  ngOnInit() {
    this.oktaAuth.getUser().then(data => this.userEmail = data['email']);

    if (this.regex.test(this.router.url)) {
      this.post = {
        id: null,
        title: '',
        body: '',
        description: '',
        user_email: ''
      };
    } else {
      this.route.params
        .subscribe(
          params => this.getPost(params['id']),
          err => console.log(err)
        );
    }
  }

  private getPost(postId: number) {
    this.postService.getPost(postId)
        .subscribe(data => this.post = data['post']);
  }

  private goBack() {
    this.location.back();
  }

  private savePost() {
    if (this.post.id === null) {
      const post = {
        title: this.post.title,
        body: this.post.body,
        description: this.post.description,
        user_email: this.userEmail
      };

      this.postService.addPost(post)
        .subscribe(
          data => console.log(data),
          err => console.log(err),
          () => this.router.navigate(['/posts'])
        );
    } else {
      this.postService.savePost(this.post)
        .subscribe(
          data => console.log(data),
          err => console.log(err),
          () => this.router.navigate(['/posts'])
        );
    }
  }

}
