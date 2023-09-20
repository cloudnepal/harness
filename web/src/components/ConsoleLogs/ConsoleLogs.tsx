import { FlexExpander, Layout, Text } from '@harnessio/uicore'
import React, { FC } from 'react'
import type { LivelogLine } from 'services/code'
import css from './ConsoleLogs.module.scss'

interface ConsoleLogsProps {
  logs: LivelogLine[]
}

export const createStreamedLogLineElement = (log: LivelogLine) => {
  const lineElement = document.createElement('div')
  lineElement.className = css.logLayout

  if (typeof log.pos === 'number') {
    const lineNumberElement = document.createElement('span')
    lineNumberElement.className = css.lineNumber
    lineNumberElement.textContent = (log.pos + 1).toString()
    lineElement.appendChild(lineNumberElement)
  }

  const logTextElement = document.createElement('span')
  logTextElement.className = css.log
  logTextElement.textContent = log.out as string
  lineElement.appendChild(logTextElement)

  const flexExpanderElement = document.createElement('span')
  flexExpanderElement.className = css.flexExpand
  lineElement.appendChild(flexExpanderElement)

  const timeElement = document.createElement('span')
  timeElement.className = css.time
  timeElement.textContent = `${log.time}s`
  lineElement.appendChild(timeElement)

  return lineElement
}

const ConsoleLogs: FC<ConsoleLogsProps> = ({ logs }) => {
  return (
    <>
      {logs.map((log, index) => {
        return (
          <Layout.Horizontal key={index} spacing={'small'} className={css.logLayout}>
            {typeof log.pos === 'number' && <Text className={css.lineNumber}>{log.pos + 1}</Text>}
            <Text className={css.log}>{log.out}</Text>
            <FlexExpander />
            <Text className={css.time}>{log.time}s</Text>
          </Layout.Horizontal>
        )
      })}
    </>
  )
}

export default ConsoleLogs
