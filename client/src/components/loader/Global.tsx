interface Props {
  highlight: string
}

function GlobalLoader({ highlight }: Props) {
  return (
    <div>{`${highlight}...`}</div>
  )
}

export default GlobalLoader
