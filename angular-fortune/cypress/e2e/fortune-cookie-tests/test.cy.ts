
describe('Webapp initialization check', () => {
    it(`Should have the title 'Fortune Cookie'`, () => {
        cy.visit('localhost:4200')
        cy.contains('Fortune Cookie');
    });
});

describe('Logout functionality check', () => {
    it(`Should have the title 'Fortune Cookie'`, () => {
        cy.visit('localhost:4200')
        cy.contains('Fortune Cookie');
    });
});