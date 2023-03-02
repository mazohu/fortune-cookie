import { HomeComponent } from './home.component'

describe('HomeComponent', () => {
  it('mounts', () => {
    cy.mount(HomeComponent)
  })

  it('default to 0', () => {
    cy.mount(HomeComponent)
    cy.get('span').should('have.text', 'Show me my fortunes!')
  })
})