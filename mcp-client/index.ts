import readline from 'readline/promises'
import dotenv from 'dotenv'
import { ChatAnthropic } from '@langchain/anthropic'
import { MCPAgent, MCPClient } from 'mcp-use'

dotenv.config({ path: '../.env' })

const ANTHROPIC_API_KEY = process.env.ANTHROPIC_API_KEY
if (!ANTHROPIC_API_KEY) {
  throw new Error('ANTHROPIC_API_KEY is not set')
}

class MCPClientClass {
  private mcp: MCPAgent
  private llm: ChatAnthropic

  constructor() {
    this.llm = new ChatAnthropic({
      apiKey: ANTHROPIC_API_KEY,
      // model: "claude-4-sonnet-20250514",
      model: 'claude-3-5-haiku-20241022',
      maxTokens: 1000,
    })

    const command = '../mcp-server/build/mcp-server'
    const config = {
      mcpServers: {
        playwright: { command, args: [] },
      },
    }
    const client = MCPClient.fromDict(config)

    this.mcp = new MCPAgent({
      llm: this.llm,
      client,
      maxSteps: 20,
    })
  }

  async processQuery(query: string) {
    const response = await this.mcp.run(query)

    return JSON.parse(JSON.stringify(response, null, 2))[0].text
  }

  async chatLoop() {
    const rl = readline.createInterface({
      input: process.stdin,
      output: process.stdout,
    })

    try {
      console.log('\nMCP Client Started!')
      console.log("Type your queries or 'quit' to exit.")

      while (true) {
        const message = await rl.question('\nQuery: ')
        if (message.toLowerCase() === 'quit') {
          break
        }

        const response = await this.processQuery(message)
        console.log('\n' + response)
      }
    } finally {
      rl.close()
    }
  }

  async cleanup() {
    await this.mcp.close()
  }
}

async function main() {
  const mcpClient = new MCPClientClass()
  try {
    await mcpClient.chatLoop()
  } finally {
    await mcpClient.cleanup()
    process.exit(0)
  }
}

main()
