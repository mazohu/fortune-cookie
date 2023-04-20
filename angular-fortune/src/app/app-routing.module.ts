import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './views/login/login.component';
import { UserpageComponent } from './views/userpage/userpage.component';
import { HomepageComponent } from './views/homepage/homepage.component';
import { EatcookieComponent } from './views/eatcookie/eatcookie.component';
import { PastfortunesComponent } from './views/pastfortunes/pastfortunes.component';
import { UserprofileComponent } from './views/userprofile/userprofile.component';

const routes: Routes = [
  { path: '', component: HomepageComponent, title: 'Fortune Cookie' }, //this is homepage
  { path: 'login', component: LoginComponent, title: 'Fortune Cookie' },
  { path: 'userpage', component: UserpageComponent, title: 'Fortune Cookie' },
  { path: 'eat-cookie', component:EatcookieComponent, title: 'Fortune Cookie'},
  { path: 'pastFortunes', component:PastfortunesComponent, title: 'Fortune Cookie'},
  { path: 'userprofile', component:UserprofileComponent, title: 'Fortune Cookie'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
//This will automatically import all the components we rout to. All we need to do is update this list
export const routingComponents = [ HomepageComponent, LoginComponent, UserpageComponent] 
