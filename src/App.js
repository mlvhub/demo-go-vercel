import React from 'react'
import { useState } from 'react'
import { formatDate } from './util'
import { useQuery } from 'react-query'
import {
  useUsersQueries, 
  fetchUsers,
  fetchUserById,
} from './api'

function App() {
  const {
    updateMutation,
    deleteMutation,
    createMutation,
  } = useUsersQueries()
  const { data: users = [], isLoading } = useQuery('users-list', fetchUsers, {
    refetchOnWindowFocus: false,
  })
  const [user, setUser] = useState({
    name: '',
    email: '',
  })

  async function handleGreet(id) {
    const user = await fetchUserById(id)
    alert(JSON.stringify(user))
  }

  async function handleDelete(id) {
    deleteMutation.mutate(id)
  }
  async function handleUpdate(id) {
    updateMutation.mutate(id)
  }

  function handleChange(event) {
    const { name, value } = event.target
    setUser(state => ({
        ...state,
        [name]: value,
    }))
  }

  async function handleSubmit(event) {
    event.preventDefault()
    createMutation.mutate(user, {
      onSuccess: () => {
        setUser({ name: '', email: '' })
      }
    })
  }

  return (
    <main className="mx-auto max-w-4x1 px-4 py-8">
      {/* TODO: https://final-form.org/docs/react-final-form/getting-started */}
      <form onSubmit={handleSubmit} className="max-w-lg mx-auto bg-white p-8 shadow-lg mb-16">
        <h1 className="text-center text-3x1 mb-8 uppercase tracking-wide font-bold text-blue-900">
          Create User
        </h1>
        <div className="mb-8">
          <label className="block mb-2 font-bold text-lg" htmlFor="email">
            Email:
          </label>
          <input onChange={handleChange} type="email" name="email" id="email" className="py-2 px-4 w-full border border-gray-400 rounded-lg" placeholder="email@email.com" />
        </div>
        <div className="mb-8">
          <label className="block mb-2 font-bold text-lg" htmlFor="name">
            Name:
          </label>
          <input onChange={handleChange} type="text" name="name" id="email" className="py-2 px-4 w-full border border-gray-400 rounded-lg" placeholder="John Doe" />
        </div>
        <button className="bg-blue-400 text-white font-bold tracking-wide uppercase py-2 px-8 rounded-lg w-full hover:shadow-lg hover:bg-blue-500 transition">
          Submit
        </button>
      </form>

      {isLoading && (
        <p className="text-center">Loading user list...</p>
      )} 

      {users.length > 0 ? <table className="w-full bg-white p-4">
        <thead>
          <tr>
            <th className="p-2 border border-gray-200 uppercase">Name</th>
            <th className="p-2 border border-gray-200 uppercase">Email</th>
            <th className="p-2 border border-gray-200 uppercase">Date</th>
            <th className="p-2 border border-gray-200 uppercase">Actions</th>
          </tr>
        </thead>
        <tbody>
          {users.map(user => (
            <tr key={user.id}>
              <td className="p-2 border border-gray-200">{user.email}</td>
              <td className="p-2 border border-gray-200">{user.name}</td>
              <td className="p-2 border border-gray-200">
                {formatDate(user.created_at)}
              </td>
              <td className="p-2 border border-gray-200">
                <div className="flex justify-between flex-col md:flex-row">
                  <button onClick={() => handleGreet(user.id)}>
                    <span role="img" aria-label="add details">
                      👀
                    </span>
                  </button>
                  <button onClick={() => handleDelete(user.id)}>
                    <span role="img" aria-label="delete user">
                     🗑 
                    </span>
                  </button>
                  <button onClick={() => handleUpdate(user.id)}>
                    <span role="img" aria-label="edit user">
                     ✏️ 
                    </span>
                  </button>
                </div>
              </td>
            </tr>
          ))}
        </tbody>
      </table> : null}
    </main>
  )
}

export default App
