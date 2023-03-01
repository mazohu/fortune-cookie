import { Component, OnInit } from '@angular/core';
import { FormBuilder } from '@angular/forms';

@Component({
  selector: 'app-eat-cookie',
  templateUrl: './eat-cookie.component.html',
  styleUrls: ['./eat-cookie.component.css']
})

export class EatCookieComponent implements OnInit {

  writeFortune = this.formBuilder.group({
    Fortune:''
  });

  constructor(
    private formBuilder: FormBuilder,

  ) {}

  ngOnInit(): void {
  }

  onSubmit(): void {
    // Process checkout data here
    console.warn('Your fortune has been submitted', this.writeFortune.value);
    this.writeFortune.reset();
  }

}
