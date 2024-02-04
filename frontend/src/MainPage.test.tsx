import { render, screen } from '@testing-library/react';
import MainPage from './MainPage';

test("MainPage가 존재함",()=>{
  render(<MainPage/>)
  const mainpage = screen.getByTitle("MainPage")
  expect(mainpage).toBeInTheDocument();
})

test("NavigationBar가 MainPage에 존재한다.", ()=>{
  render(<MainPage/>)
  const navigationbar = screen.getByTitle("NavigationBar")
  expect(navigationbar).toBeInTheDocument();
})

test("MainBody가 MainPage에 존재한다.", ()=>{
  render(<MainPage/>)
  const mainbody = screen.getByTitle("MainBody")
  expect(mainbody).toBeInTheDocument();
})