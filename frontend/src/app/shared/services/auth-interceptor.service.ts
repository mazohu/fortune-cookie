import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor
} from '@angular/common/http';
import { Observable, from } from 'rxjs';
import { OktaAuthStateService } from '@okta/okta-angular';
import { OktaAuth } from '@okta/okta-auth-js'

@Injectable()
export class AuthInterceptor implements HttpInterceptor {

  constructor(private oktaAuth: OktaAuth) {
  }

  intercept(request: HttpRequest<any>, next: HttpHandler): Observable<any> {
    return from(this.handleAccess(request, next));
  }

  private async handleAccess(request: HttpRequest<any>, next: HttpHandler): Promise<any> {
      const accessToken = await this.oktaAuth.getAccessToken();
      request = request.clone({
        setHeaders: {
          Authorization: `Bearer ${accessToken}`
        }
      });
      return next.handle(request).toPromise();
  }
}