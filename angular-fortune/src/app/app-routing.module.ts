import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './views/login/login.component';
import { UserpageComponent } from './views/userpage/userpage.component';
import { HomepageComponent } from './views/homepage/homepage.component';
import { EatcookieComponent } from './views/eatcookie/eatcookie.component';
import { PastfortunesComponent } from './views/pastfortunes/pastfortunes.component';

const routes: Routes = [
  { path: '', component: HomepageComponent }, //this is homepage
  { path: 'login', component: LoginComponent },
  { path: 'userpage', component: UserpageComponent },
  { path: 'eat-cookie', component:EatcookieComponent},
  { path: 'pastFortunes', component:PastfortunesComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
//This will automatically import all the components we rout to. All we need to do is update this list
export const routingComponents = [ HomepageComponent, LoginComponent, UserpageComponent] 
