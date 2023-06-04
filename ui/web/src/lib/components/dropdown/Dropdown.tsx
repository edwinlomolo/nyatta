import { type MouseEventHandler } from 'react'
import { Menu, MenuList, MenuButton, Portal, MenuItem } from '@chakra-ui/react'

interface Option {
  text: React.ReactNode
  onClick?: MouseEventHandler<HTMLButtonElement>
}

interface Props {
  children: React.ReactNode
  options: Option[]
}

function Dropdown ({ children, options }: Props) {
  return (
    <Menu isLazy>
     <MenuButton type="button">
       {children}
     </MenuButton>
     <Portal>
       <MenuList>
        {options.map((item: Option, index: number) => (
          <MenuItem
            key={index}
            onClick={(item.onClick != null) ? item.onClick : undefined}
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
