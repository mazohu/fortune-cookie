
describe('Webapp initialization check', () => {
    it(`Should have the title 'Fortune Cookie'`, () => {
        cy.visit('localhost:4200'); 
        cy.contains('Fortune Cookie');
    });
});

describe('Login functionality check', () => {
    it(`Should bring you to the authentication page`, () => {
        cy.visit('localhost:4200');
        cy.get('[data-test=login-check]').click();
        cy.contains('\"G\"');
    });
});

describe('Non-signed in user cannot go to /userpage', () => {
    it(`Should prompt the Google sign in`, () => {
        cy.visit('localhost:4200/userpage');
        cy.contains('\"G\"');
    });
});

describe('Non-signed in user cannot go to /userprofile', () => {
    it(`Should prompt the Google sign in`, () => {
        cy.visit('localhost:4200/userprofile');
        cy.contains('\"G\"');
    });
});

describe('Non-signed in user cannot go to /eat-cookie', () => {
    it(`Should prompt the Google sign in`, () => {
        cy.visit('localhost:4200/eat-cookie');
        cy.contains('\"G\"');
    });
});

describe('Non-signed in user cannot go to /pastFortunes', () => {
    it(`Should prompt the Google sign in`, () => {
        cy.visit('localhost:4200/pastFortunes');
        cy.contains('\"G\"');
    });
});