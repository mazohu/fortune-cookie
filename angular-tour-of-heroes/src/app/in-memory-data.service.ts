import { Injectable } from '@angular/core';
import { InMemoryDbService } from 'angular-in-memory-web-api';
import { Hero } from './hero';

@Injectable({
  providedIn: 'root'
})
export class InMemoryDataService implements InMemoryDbService {

  createDb() {
    const heroes = [
      { id: 12, name: 'You will have a good day' },
      { id: 13, name: 'You will have a bad day' },
      { id: 14, name: 'You are what you eat' },
      { id: 15, name: 'Easy peasy lemon squeezy, difficult shmifficult difficult difficult' },
      { id: 16, name: 'Don\'t kill the part of you that is cringe. Kill the part of you that cringes.' },
      { id: 17, name: 'You will go crazy' },
      { id: 18, name: 'You will get a job as an Angular web developer' },
      { id: 19, name: 'You will get a job as a Golang developer' },
      { id: 20, name: 'A hole will open up in the center of the earth and swallow you so that you don\'t have to worry about this class' }
    ];
    return {heroes};
  }

  // overrides the genId method to ensure that a hero always has an id.
  // If the heroes array is empty, the method below returns the initial number (11).
  // If the heroes array is not empty, the method below returns the highest hero id + 1.
  genId(heroes: Hero[]): number {
    return heroes.length > 0 ? Math.max(...heroes.map(hero => hero.id)) + 1 : 11;
  }
}
