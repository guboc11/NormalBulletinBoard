import { Box, Typography } from "@mui/material"
import NavigationBar from "./NavigationBar"
import MainBody from "./MainBody"

export default function MainPage() {
  return(
    <Box title="MainPage">
      <Typography>Main Page</Typography>
      <NavigationBar></NavigationBar>
      <MainBody></MainBody>
    </Box>
  )
}