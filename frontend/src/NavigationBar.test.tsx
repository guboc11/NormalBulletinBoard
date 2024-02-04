import { render, screen } from '@testing-library/react';
import NavigationBar from './NavigationBar';

test("NavigationBar가 존재한다.", ()=>{
  render(<NavigationBar/>)
  const navigationbar = screen.getByTitle("NavigationBar")
  expect(navigationbar).toBeInTheDocument();
})