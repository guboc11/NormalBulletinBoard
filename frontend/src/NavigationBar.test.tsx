import { render, screen } from '@testing-library/react';
import NavigationBar from './NavigationBar';

test("NavigationBar가 존재한다.", ()=>{
  render(<NavigationBar/>)
  const navigationbar = screen.getByTitle("NavigationBar")
  expect(navigationbar).toBeInTheDocument();
})

test("ViewAll 버튼, NewPost 버튼이 존재한다.", ()=>{
  render(<NavigationBar/>)
  const viewAll = screen.getByTitle("ViewAllButton")
  const newPost = screen.getByTitle("NewPostButton")
  expect(viewAll).toBeInTheDocument();
  expect(newPost).toBeInTheDocument();
})