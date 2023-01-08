import { MouseEventHandler } from 'react'
import { Menu, MenuList, MenuButton, Portal, MenuItem } from '@chakra-ui/react'

interface Option {
  text: string
  onClick?: MouseEventHandler<HTMLButtonElement>
}

interface Props {
  children: React.ReactNode
  options: Option[]
}

function Dropdown({ children, options }: Props) {
  return (
    <Menu>
     <MenuButton type="button">
       {children}
     </MenuButton>
     <Portal>
       <MenuList>
        {options.map((item: Option, index: number) => (
          <MenuItem
            key={index}
            onClick={item.onClick ? item.onClick : undefined}
          >
            {item.text}
          </MenuItem>
        ))}
       </MenuList>
     </Portal>
    </Menu>
  )
}

export default Dropdown
